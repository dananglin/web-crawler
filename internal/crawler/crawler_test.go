package crawler

import (
	"fmt"
	"slices"
	"testing"

	"codeflow.dananglin.me.uk/apollo/web-crawler/internal/util"
)

func TestCrawler(t *testing.T) {
	testBaseURL := "https://example.com"

	testCrawler, err := NewCrawler(testBaseURL, 1, 10)
	if err != nil {
		t.Fatalf("Test 'TestCrawler' FAILED: unexpected error creating the crawler: %v", err)
	}

	testCasesForEqualDomains := []struct {
		name   string
		rawURL string
		want   bool
	}{
		{
			name:   "Same domain",
			rawURL: "https://example.com",
			want:   true,
		},
		{
			name:   "Same domain, different path",
			rawURL: "https://example.com/about/contact",
			want:   true,
		},
		{
			name:   "Same domain, different protocol",
			rawURL: "http://example.com",
			want:   true,
		},
		{
			name:   "Different domain",
			rawURL: "https://blog.person.me.uk",
			want:   false,
		},
		{
			name:   "Different domain, same path",
			rawURL: "https://example.org/blog",
			want:   false,
		},
	}

	for ind, tc := range slices.All(testCasesForEqualDomains) {
		t.Run(tc.name, testIsInternalLink(
			testCrawler,
			ind+1,
			tc.name,
			tc.rawURL,
			tc.want,
		))
	}

	testCasesForPages := []struct {
		rawURL      string
		wantVisited bool
	}{
		{
			rawURL:      "https://example.com/tags/linux",
			wantVisited: false,
		},
		{
			rawURL:      "https://example.com/blog",
			wantVisited: false,
		},
		{
			rawURL:      "https://example.com/about/contact.html",
			wantVisited: false,
		},
		{
			rawURL:      "https://example.com/blog",
			wantVisited: true,
		},
	}

	for ind, tc := range slices.All(testCasesForPages) {
		name := fmt.Sprintf("Adding %s to the pages map", tc.rawURL)
		t.Run(name, testHasVisited(
			testCrawler,
			ind+1,
			name,
			tc.rawURL,
			tc.wantVisited,
		))
	}
}

func testIsInternalLink(
	testCrawler *Crawler,
	testNum int,
	testName string,
	rawURL string,
	want bool,
) func(t *testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		got, err := testCrawler.isInternalLink(rawURL)
		if err != nil {
			t.Fatalf(
				"Test %d - '%s' FAILED: unexpected error: %v",
				testNum,
				testName,
				err,
			)
		}

		if got != want {
			t.Errorf(
				"Test %d - '%s' FAILED: unexpected domain comparison received: want %t, got %t",
				testNum,
				testName,
				want,
				got,
			)
		} else {
			t.Logf(
				"Test %d - '%s' PASSED: expected domain comparison received: got %t",
				testNum,
				testName,
				got,
			)
		}
	}
}

func testHasVisited(
	testCrawler *Crawler,
	testNum int,
	testName string,
	rawURL string,
	wantVisited bool,
) func(t *testing.T) {
	return func(t *testing.T) {
		normalisedURL, err := util.NormaliseURL(rawURL)
		if err != nil {
			t.Fatalf(
				"Test %d - '%s' FAILED: unexpected error: %v",
				testNum,
				testName,
				err,
			)
		}

		gotVisited := testCrawler.addPageVisit(normalisedURL, true)

		if gotVisited != wantVisited {
			t.Errorf(
				"Test %d - '%s' FAILED: unexpected bool returned after updated pages record: want %t, got %t",
				testNum,
				testName,
				wantVisited,
				gotVisited,
			)
		} else {
			t.Logf(
				"Test %d - '%s' PASSED: expected bool returned after updated pages record: got %t",
				testNum,
				testName,
				gotVisited,
			)
		}
	}
}
