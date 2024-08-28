package crawler

import (
	"cmp"
	"maps"
	"slices"
	"strconv"
	"strings"
)

type report struct {
	Format  string   `json:"-"`
	BaseURL string   `json:"baseUrl"`
	Records []record `json:"records"`
}

type record struct {
	Link     string `json:"link"`
	Count    int    `json:"count"`
	LinkType string `json:"linkType"`
}

func newReport(format, baseURL string, pages map[string]pageStat) report {
	records := make([]record, 0)

	for link, stats := range maps.All(pages) {
		linkType := "internal"
		if !stats.internal {
			linkType = "external"
		}

		record := record{
			Link:     link,
			Count:    stats.count,
			LinkType: linkType,
		}

		records = append(records, record)
	}

	report := report{
		Format:  format,
		BaseURL: baseURL,
		Records: records,
	}

	report.sortRecords()

	return report
}

func (r *report) sortRecords() {
	// First sort records by count (in reverse order hopefully)
	// Then sort records by name if two elements have the same count.
	slices.SortFunc(r.Records, func(a, b record) int {
		if n := cmp.Compare(a.Count, b.Count); n != 0 {
			return -1 * n
		}

		return strings.Compare(a.Link, b.Link)
	})
}

func (r report) String() string {
	switch r.Format {
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
	builder.WriteString("\n" + "REPORT for " + r.BaseURL)
	builder.WriteString("\n" + titlebar)

	for ind := range slices.All(r.Records) {
		links := "links"
		if r.Records[ind].Count == 1 {
			links = "link"
		}

		builder.WriteString("\nFound " + strconv.Itoa(r.Records[ind].Count) + " " + r.Records[ind].LinkType + " " + links + " to " + r.Records[ind].Link)
	}

	return builder.String()
}

func (r report) csv() string {
	var builder strings.Builder

	builder.WriteString("LINK,TYPE,COUNT")

	for ind := range slices.All(r.Records) {
		builder.WriteString("\n" + r.Records[ind].Link + "," + r.Records[ind].LinkType + "," + strconv.Itoa(r.Records[ind].Count))
	}

	return builder.String()
}
