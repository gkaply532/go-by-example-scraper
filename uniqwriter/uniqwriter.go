package uniqwriter

import (
	"bufio"
	"fmt"
	"github.com/gkaply532/go-by-example-scraper/set"
	"io"
)

func New(wrapped io.Writer) io.Writer {
	reader, writer := io.Pipe()

	go func() {
		scanner := bufio.NewScanner(reader)
		written := set.New[string]()
		for scanner.Scan() {
			line := scanner.Text()
			if !written.Present(line) {
				fmt.Fprintln(wrapped, line)
				written.Add(line)
			}
		}
	}()

	return writer
}
