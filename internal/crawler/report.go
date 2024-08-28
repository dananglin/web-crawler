package crawler

import (
	"cmp"
	"maps"
	"slices"
	"strconv"
	"strings"
)

type report struct {
	baseURL string
	records []record
}

type record struct {
	link     string
	count    int
	internal bool
}

func newReport(baseURL string, pages map[string]pageStat) report {
	records := make([]record, 0)

	for link, stats := range maps.All(pages) {
		record := record{
			link:     link,
			count:    stats.count,
			internal: stats.internal,
		}

		records = append(records, record)
	}

	report := report{
		baseURL: baseURL,
		records: records,
	}

	report.sortRecords()

	return report
}

func (r *report) sortRecords() {
	// First sort records by count (in reverse order hopefully)
	// Then sort records by name if two elements have the same count.
	slices.SortFunc(r.records, func(a, b record) int {
		if n := cmp.Compare(a.count, b.count); n != 0 {
			return -1 * n
		}

		return strings.Compare(a.link, b.link)
	})
}

func (r report) String() string {
	var builder strings.Builder

	titlebar := strings.Repeat("\u2500", 80)

	builder.WriteString("\n" + titlebar)
	builder.WriteString("\n" + "REPORT for " + r.baseURL)
	builder.WriteString("\n" + titlebar)

	for ind := range slices.All(r.records) {
		linkType := "internal"
		if !r.records[ind].internal {
			linkType = "external"
		}

		links := "links"
		if r.records[ind].count == 1 {
			links = "link"
		}

		builder.WriteString("\nFound " + strconv.Itoa(r.records[ind].count) + " " + linkType + " " + links + " to " + r.records[ind].link)
	}

	builder.WriteString("\n")

	return builder.String()
}
