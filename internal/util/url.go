package util

import (
	"fmt"
	"net/url"
	"strings"
)

func NormaliseURL(rawURL string) (string, error) {
	const normalisedFormat string = "%s%s"

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("error parsing the URL %q: %w", rawURL, err)
	}

	return fmt.Sprintf(normalisedFormat, parsedURL.Hostname(), strings.TrimSuffix(parsedURL.Path, "/")), nil
}
