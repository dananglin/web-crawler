package util_test

import (
	"os"
	"reflect"
	"slices"
	"testing"

	"codeflow.dananglin.me.uk/apollo/web-crawler/internal/util"
)

func TestGetURLsFromHTML(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		filepath string
		baseURL  string
		want     []string
	}{
		{
			name:     "HTML documentation using blog.boot.dev",
			filepath: "testdata/GetURLFromHTML/blog.boot.dev.html",
			baseURL:  "https://blog.boot.dev",
			want: []string{
				"https://blog.boot.dev/path/one",
				"https://other.com/path/one",
			},
		},
		{
			name:     "HTML documentation using https://ben-bartlett.me.uk",
			filepath: "testdata/GetURLFromHTML/ben-bartlett.html",
			baseURL:  "https://ben-bartlett.me.uk",
			want: []string{
				"https://ben-bartlett.me.uk",
				"https://github.com/ben-bartlett",
				"https://mastodon.ben-bartlett.me.uk",
				"https://ben-bartlett.me.uk/blog",
				"https://ben-bartlett.me.uk/projects/orange-juice",
				"https://ben-bartlett.me.uk/projects/mustangs",
				"https://ben-bartlett.me.uk/projects/honeycombs",
			},
		},
		{
			name:     "HTML documentation using https://simple.cooking",
			filepath: "testdata/GetURLFromHTML/my-simple-cooking-website.html",
			baseURL:  "https://simple.cooking",
			want: []string{
				"https://simple.cooking/recipes/sweet-n-sour-kung-pao-style-chicken",
				"https://simple.cooking/recipes/beef-and-broccoli",
				"https://simple.cooking/recipes/asian-glazed-salmon",
				"https://simple.cooking/recipes/caesar-salad",
				"https://simple.cooking/recipes/simple-tuna-salad",
				"https://simple.cooking/recipes/wholemeal-pizza",
				"https://simple.cooking/news",
				"https://simple.cooking/about/contact",
				"https://the-other-site.example.new/home",
			},
		},
	}

	for _, tc := range slices.All(cases) {
		t.Run(tc.name, testGetURLsFromHTML(tc.filepath, tc.baseURL, tc.want))
	}
}

func testGetURLsFromHTML(path, baseURL string, want []string) func(t *testing.T) {
	failedTestPrefix := "Test TestGetURLsFromHTML FAILED:"

	return func(t *testing.T) {
		t.Parallel()

		htmlDoc, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("%s unable to open read data from %s: %v", failedTestPrefix, path, err)
		}

		got, err := util.GetURLsFromHTML(string(htmlDoc), baseURL)
		if err != nil {
			t.Fatalf(
				"Test TestGetURLsFromHTML FAILED: unexpected error: %v",
				err,
			)
		}

		if !reflect.DeepEqual(want, got) {
			t.Errorf(
				"Test TestGetURLsFromHTML FAILED: unexpected URLs found in HTML body: want %v, got %v",
				want,
				got,
			)
		} else {
			t.Logf(
				"Test TestGetURLsFromHTML PASSED: expected URLs found in HTML body: got %v",
				got,
			)
		}
	}
}
