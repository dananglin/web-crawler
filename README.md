# Web Crawler

## Overview

This web crawler crawls a given URL and generates a report for all the internal links it finds.

### Repository mirrors

- **Code Flow:** https://codeflow.dananglin.me.uk/apollo/web-crawler
- **GitHub:** https://github.com/dananglin/web-crawler

## Requirements

- **Go:** A minimum version of Go 1.23.0 is required for building/installing the web crawler. Please go [here](https://go.dev/dl/) to download the latest version.

## How to run the application

Clone this repository to your local machine.
```
git clone https://github.com/dananglin/web-crawler.git
```

Build the application.
```
go build -o crawler .
```

Run the application specifying the website that you want to crawl.

- To crawl `https://example.com` using 3 concurrent workers and generate a report of up to 20 unique discovered pages:
   ```
   ./crawler --max-workers 3 --max-pages 20 https://example.com
   ```

## Flags

You can configure the application with the following flags.

| Name | Description | Default |
|------|-------------|---------|
| `max-workers` | The maximum number of concurrent workers. | 2 |
| `max-pages` | The maximum number of pages discovered before stopping the crawl. | 10 |
