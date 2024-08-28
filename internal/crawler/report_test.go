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
		Format:  "text",
		BaseURL: "https://example.org",
		Records: []record{
			{Link: "example.org", Count: 45, LinkType: "internal"},
			{Link: "example.org/about/contact", Count: 10, LinkType: "internal"},
			{Link: "example.org/posts", Count: 4, LinkType: "internal"},
			{Link: "example.org/tags", Count: 4, LinkType: "internal"},
			{Link: "mastodon.example.social/@benbarlett", Count: 4, LinkType: "external"},
			{Link: "example.org/tags/golang", Count: 2, LinkType: "internal"},
			{Link: "ben-barlett.dev", Count: 1, LinkType: "external"},
			{Link: "example.org/posts/yet-another-web-crawler-has-emerged", Count: 1, LinkType: "internal"},
			{Link: "github.com/benbarlettdotdev", Count: 1, LinkType: "external"},
			{Link: "github.com/dananglin/web-crawler", Count: 1, LinkType: "external"},
		},
	}

	got := newReport(format, testBaseURL, testPages)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Test 'TestReport' FAILED: unexpected report created, want: %v\n\nbut got: %v", want, got)
	} else {
		t.Logf("Test 'TestReport' PASSED: expected report created, got: %v", got)
	}
}
