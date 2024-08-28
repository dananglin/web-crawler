package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"codeflow.dananglin.me.uk/apollo/web-crawler/internal/crawler"
)

func main() {
	if err := run(); err != nil {
		os.Stderr.WriteString("ERROR: " + err.Error() + "\n")

		os.Exit(1)
	}
}

var errNoURLProvided = errors.New("the URL is not provided")

func run() error {
	var (
		maxWorkers int
		maxPages   int
		format     string
		file       string
	)

	flag.IntVar(&maxWorkers, "max-workers", 2, "The maximum number of concurrent workers")
	flag.IntVar(&maxPages, "max-pages", 10, "The maximum number of pages to discover before stopping the crawl")
	flag.StringVar(&format, "format", "text", "The format of the report. Valid formats are 'text', 'json' and 'csv'")
	flag.StringVar(&file, "file", "", "The file to save the report to")

	flag.Parse()

	if flag.NArg() < 1 {
		return errNoURLProvided
	}

	baseURL := flag.Arg(0)

	c, err := crawler.NewCrawler(baseURL, maxWorkers, maxPages, format, file)
	if err != nil {
		return fmt.Errorf("unable to create the crawler: %w", err)
	}

	go c.Crawl(baseURL)

	c.Wait()

	if err := c.GenerateReport(); err != nil {
		return fmt.Errorf("unable to generate the report: %w", err)
	}

	return nil
}
