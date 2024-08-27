package util

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func GetURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	htmlDoc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, fmt.Errorf("unable to parse the HTML document: %w", err)
	}

	parsedRawBaseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return []string{}, fmt.Errorf("unable to parse the raw base URL %q: %w", rawBaseURL, err)
	}

	output := make([]string, 0, 3)

	var extractLinkFunc func(*html.Node) error

	extractLinkFunc = func(node *html.Node) error {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, a := range node.Attr {
				if a.Key == "href" {
					extractedURL, err := getAbsoluteURL(a.Val, parsedRawBaseURL)
					if err != nil {
						return fmt.Errorf("unable to get the absolute URL of %s: %w", a.Val, err)
					}

					output = append(output, extractedURL)

					break
				}
			}
		}

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			if err := extractLinkFunc(c); err != nil {
				return err
			}
		}

		return nil
	}

	if err := extractLinkFunc(htmlDoc); err != nil {
		return []string{}, err
	}

	return output, nil
}

func getAbsoluteURL(inputURL string, baseURL *url.URL) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", fmt.Errorf("unable to parse the URL from %s: %w", inputURL, err)
	}

	if parsedURL.Scheme == "" && parsedURL.Host == "" {
		parsedURL.Scheme = baseURL.Scheme
		parsedURL.Host = baseURL.Host
	}

	return parsedURL.String(), nil
}
