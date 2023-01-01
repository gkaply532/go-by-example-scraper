# Go by Example Scraper

After going through [A Tour of Go](https://go.dev/tour/), doing the given
exercises and then reading [Go by Example](https://gobyexample.com/), I decided
that my first go project will be a web scraper that finds links given in the go
by example website so I can read the mentioned blog posts, find out the used
packages etc.

## TODO
 - [x] don't scrape if not under `gobyexample.com`
 - [x] scrape pages once
 - [x] output links that are not under `gobyexample.com`
 - [x] output links once (like uniq)
 - [ ] scrape in parallel
 - [ ] implement rate limiting
 - [ ] output which link was under which page
 - [ ] output a list of imported modules for each page
 - [ ] sniff and verify "Content-Type" of requests.
 - [ ] implement scraping by sub page instead of origin
 - [ ] fix same origin check
