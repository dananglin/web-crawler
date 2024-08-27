package crawler

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func getHTML(rawURL string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return "", fmt.Errorf("error creating the HTTP request: %w", err)
	}

	client := http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		return "", fmt.Errorf("error getting the response: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf(
			"received a bad status from %s: (%d) %s",
			rawURL,
			resp.StatusCode,
			resp.Status,
		)
	}

	contentType := resp.Header.Get("content-type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("unexpected content type received: want text/html, got %s", contentType)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading the data from the response: %w", err)
	}

	return string(data), nil
}
