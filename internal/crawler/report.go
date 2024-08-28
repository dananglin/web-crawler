package crawler

import (
	"cmp"
	"maps"
	"slices"
	"strconv"
	"strings"
)

type report struct {
	format  string
	baseURL string
	records []record
}

type record struct {
	link     string
	count    int
	linkType string
}

func newReport(format, baseURL string, pages map[string]pageStat) report {
	records := make([]record, 0)

	for link, stats := range maps.All(pages) {
		linkType := "internal"
		if !stats.internal {
			linkType = "external"
		}

		record := record{
			link:     link,
			count:    stats.count,
			linkType: linkType,
		}

		records = append(records, record)
	}

	report := report{
		format:  format,
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
	switch r.format {
	case "csv":
		return r.csv()
	default:
		return r.text()
	}
}

func (r report) text() string {
	var builder strings.Builder

	titlebar := strings.Repeat("\u2500", 80)

	builder.WriteString("\n" + titlebar)
	builder.WriteString("\n" + "REPORT for " + r.baseURL)
	builder.WriteString("\n" + titlebar)

	for ind := range slices.All(r.records) {
		links := "links"
		if r.records[ind].count == 1 {
			links = "link"
		}

		builder.WriteString("\nFound " + strconv.Itoa(r.records[ind].count) + " " + r.records[ind].linkType + " " + links + " to " + r.records[ind].link)
	}

	return builder.String()
}

func (r report) csv() string {
	var builder strings.Builder

	builder.WriteString("LINK,TYPE,COUNT")

	for ind := range slices.All(r.records) {
		builder.WriteString("\n" + r.records[ind].link + "," + r.records[ind].linkType + "," + strconv.Itoa(r.records[ind].count))
	}

	return builder.String()
}
