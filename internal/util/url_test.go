package util_test

import (
	"slices"
	"testing"

	"codeflow.dananglin.me.uk/apollo/web-crawler/internal/util"
)

func TestNormaliseURL(t *testing.T) {
	t.Parallel()

	wantNormalisedURL := "blog.boot.dev/path"

	cases := []struct {
		name     string
		inputURL string
	}{
		{
			name:     "remove HTTPS scheme",
			inputURL: "https://blog.boot.dev/path",
		},
		{
			name:     "remove HTTP scheme",
			inputURL: "http://blog.boot.dev/path",
		},
		{
			name:     "remove HTTPS scheme with a trailing slash",
			inputURL: "https://blog.boot.dev/path/",
		},
		{
			name:     "remove HTTP scheme with a trailing slash",
			inputURL: "http://blog.boot.dev/path/",
		},
		{
			name:     "remove HTTPS scheme with port 443",
			inputURL: "https://blog.boot.dev:443/path",
		},
		{
			name:     "remove HTTP scheme with port 80",
			inputURL: "http://blog.boot.dev:80/path",
		},
		{
			name:     "normalised URL",
			inputURL: "blog.boot.dev/path",
		},
	}

	for ind, tc := range slices.All(cases) {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := util.NormaliseURL(tc.inputURL)
			if err != nil {
				t.Fatalf(
					"Test %d - '%s' FAILED: unexpected error: %v",
					ind,
					tc.name,
					err,
				)
			}

			if got != wantNormalisedURL {
				t.Errorf(
					"Test %d - %s FAILED: unexpected normalised URL returned: want %s, got %s",
					ind,
					tc.name,
					wantNormalisedURL,
					got,
				)
			} else {
				t.Logf(
					"Test %d - %s PASSED: expected normalised URL returned: got %s",
					ind,
					tc.name,
					got,
				)
			}
		})
	}
}
