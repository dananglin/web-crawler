package crawler

import (
	"reflect"
	"testing"
)

func TestReport(t *testing.T) {
	t.Parallel()

	format := "text"
	testBaseURL := "https://example.org"
	testPages := map[string]pageStat{
		"mastodon.example.social/@benbarlett":                   {count: 4, internal: false},
		"example.org/posts/yet-another-web-crawler-has-emerged": {count: 1, internal: true},
		"example.org/about/contact":                             {count: 10, internal: true},
		"github.com/benbarlettdotdev":                           {count: 1, internal: false},
		"example.org/posts":                                     {count: 4, internal: true},
		"github.com/dananglin/web-crawler":                      {count: 1, internal: false},
		"ben-barlett.dev":                                       {count: 1, internal: false},
		"example.org":                                           {count: 45, internal: true},
		"example.org/tags":                                      {count: 4, internal: true},
		"example.org/tags/golang":                               {count: 2, internal: true},
	}

	want := report{
		format:  "text",
		baseURL: "https://example.org",
		records: []record{
			{link: "example.org", count: 45, linkType: "internal"},
			{link: "example.org/about/contact", count: 10, linkType: "internal"},
			{link: "example.org/posts", count: 4, linkType: "internal"},
			{link: "example.org/tags", count: 4, linkType: "internal"},
			{link: "mastodon.example.social/@benbarlett", count: 4, linkType: "external"},
			{link: "example.org/tags/golang", count: 2, linkType: "internal"},
			{link: "ben-barlett.dev", count: 1, linkType: "external"},
			{link: "example.org/posts/yet-another-web-crawler-has-emerged", count: 1, linkType: "internal"},
			{link: "github.com/benbarlettdotdev", count: 1, linkType: "external"},
			{link: "github.com/dananglin/web-crawler", count: 1, linkType: "external"},
		},
	}

	got := newReport(format, testBaseURL, testPages)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Test 'TestReport' FAILED: unexpected report created, want: %v\n\nbut got: %v", want, got)
	} else {
		t.Logf("Test 'TestReport' PASSED: expected report created, got: %v", got)
	}
}
