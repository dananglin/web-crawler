package main

import (
	"fmt"
	"os"
	"strconv"

	"codeflow.dananglin.me.uk/apollo/web-crawler/internal/crawler"
)

func main() {
	if err := run(); err != nil {
		os.Stderr.WriteString("ERROR: " + err.Error() + "\n")

		os.Exit(1)
	}
}

func run() error {
	args := os.Args[1:]

	if len(args) != 3 {
		return fmt.Errorf("unexpected number of arguments received: want 3, got %d", len(args))
	}

	baseURL := args[0]

	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("unable to convert the max concurrency (%s) to an integer: %w", args[1], err)
	}

	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		return fmt.Errorf("unable to convert the max pages (%s) to an integer: %w", args[2], err)
	}

	c, err := crawler.NewCrawler(baseURL, maxConcurrency, maxPages)
	if err != nil {
		return fmt.Errorf("unable to create the crawler: %w", err)
	}

	go c.Crawl(baseURL)

	c.Wait()

	c.PrintReport()

	return nil
}
