package main

import (
	"errors"
	"fmt"
	"github.com/gkaply532/go-by-example-scraper/set"
	"github.com/gkaply532/go-by-example-scraper/uniqwriter"
	"golang.org/x/net/html"
	"mime"
	"net/http"
	"net/url"
	"os"
)

func getAndParseUrl(base url.URL) (*html.Node, error) {
	resp, err := http.Get(base.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return nil, err
	}

	if mediaType != "text/html" {
		return nil, errors.New("not text/html")
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func walkNodes(root *html.Node) <-chan *html.Node {
	channel := make(chan *html.Node)
	go func() {
		node := root
		for {
			channel <- node
			if node.FirstChild != nil {
				node = node.FirstChild
			} else {
				for node.NextSibling == nil {
					if node == root {
						close(channel)
						return
					}
					node = node.Parent
				}
				node = node.NextSibling
			}
		}
	}()
	return channel
}

func grabLinks(nodes <-chan *html.Node) <-chan string {
	channel := make(chan string)
	go func() {
		for node := range nodes {
			if !(node.Type == html.ElementNode && node.Data == "a") {
				continue
			}
			for _, attr := range node.Attr {
				if attr.Key != "href" {
					continue
				}
				channel <- attr.Val
			}
		}
		close(channel)
	}()
	return channel
}

func isSameOrigin(a, b url.URL) bool {
	return a.Scheme == b.Scheme && a.Host == b.Host
}

func main() {
	site, _ := url.Parse("https://gobyexample.com/")

	pendingLinks := set.New[url.URL]()
	scrapedLinks := set.New[url.URL]()

	pendingLinks.Add(*site)

	uniqWriter := uniqwriter.New(os.Stdout)

	for pendingLinks.Len() > 0 {
		page, _ := pendingLinks.First()

		doc, err := getAndParseUrl(page)
		if err != nil {
			panic(err)
		}

		links := grabLinks(walkNodes(doc))

		for link := range links {
			if uri, err := page.Parse(link); err != nil {
				fmt.Fprintf(os.Stderr, "can't parse link %q, error: %v\n", link, err)
			} else if isSameOrigin(page, *uri) {
				if !scrapedLinks.Present(*uri) {
					pendingLinks.Add(*uri)
				}
			} else {
				fmt.Fprintln(uniqWriter, uri)
			}
		}

		pendingLinks.Remove(page)
		scrapedLinks.Add(page)
	}
}
