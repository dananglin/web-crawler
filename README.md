# Web Crawler

## Overview

This web crawler crawls a given website and generates a report for all the internal and external links found during the crawl.

### Repository mirrors

- **Code Flow:** https://codeflow.dananglin.me.uk/apollo/web-crawler
- **GitHub:** https://github.com/dananglin/web-crawler

## Requirements

- **Go:** A minimum version of Go 1.23.0 is required for building/installing the web crawler. Please go [here](https://go.dev/dl/) to download the latest version.

## Build the application

Clone this repository to your local machine.
```
git clone https://github.com/dananglin/web-crawler.git
```

Build the application.

- Build with go
   ```
   go build -o crawler .
   ```
- Or build with [mage](https://magefile.org/) if you have it installed.
   ```
   mage build
   ```

## Run the application

Run the application specifying the website that you want to crawl.

### Format

`./crawler [FLAGS] URL`

### Examples

- Crawl the [Crawler Test Site](https://crawler-test.com).
  ```
  ./crawler https://crawler-test.com
  ```
- Crawl the site using 3 concurrent workers and stop the crawl after discovering a maximum of 100 unique pages.
   ```
   ./crawler --max-workers 3 --max-pages 100 https://crawler-test.com
   ```
- Crawl the site and print out a CSV report.
   ```
   ./crawler --max-workers 3 --max-pages 100 --format csv https://crawler-test.com
   ```
- Crawl the site and save the report to a CSV file.
   ```
   mkdir -p reports
   ./crawler --max-workers 3 --max-pages 100 --format csv --file reports/report.csv https://crawler-test.com
   ```

## Flags

You can configure the application with the following flags.

| Name | Description | Default |
|------|-------------|---------|
| `max-workers` | The maximum number of concurrent workers. | 2 |
| `max-pages` | The maximum number of pages the crawler can discoverd before stopping the crawl. | 10 |
| `format` | The format of the generated report.<br>Currently supports `text` and `csv`. | text |
| `file` | The file to save the generated report to.<br>Leave this empty to print to the screen instead. | |
