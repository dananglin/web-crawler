package report

import (
	"cmp"
	"maps"
	"slices"
	"strconv"
	"strings"
)

type Report struct {
	baseURL string
	records []Record
}

type Record struct {
	link  string
	count int
}

func NewReport(baseURL string, pages map[string]int) Report {
	records := make([]Record, 0)

	for link, count := range maps.All(pages) {
		records = append(records, Record{link: link, count: count})
	}

	report := Report{
		baseURL: baseURL,
		records: records,
	}

	report.SortRecords()

	return report
}

func (r *Report) SortRecords() {
	// First sort records by count (in reverse order hopefully)
	// Then sort records by name if two elements have the same count.
	slices.SortFunc(r.records, func(a, b Record) int {
		if n := cmp.Compare(a.count, b.count); n != 0 {
			return -1 * n
		}

		return strings.Compare(a.link, b.link)
	})
}

func (r Report) String() string {
	var builder strings.Builder

	titlebar := strings.Repeat("\u2500", 80)

	builder.WriteString("\n" + titlebar)
	builder.WriteString("\n" + "REPORT for " + r.baseURL)
	builder.WriteString("\n" + titlebar)

	for ind := range slices.All(r.records) {
		builder.WriteString("\nFound " + strconv.Itoa(r.records[ind].count) + " internal links to " + r.records[ind].link)
	}

	builder.WriteString("\n")

	return builder.String()
}
