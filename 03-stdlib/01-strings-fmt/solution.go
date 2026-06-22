package stringsfmt

// Slugify converts s to a URL-friendly slug: trim space, lowercase, replace
// spaces and underscores with hyphens, collapse repeated hyphens.
func Slugify(s string) string {
	return ""
}

// FormatTable renders headers and rows as pipe-separated columns.
// Example: headers ["a","b"], rows [["1","2"]] → "a | b\n1 | 2\n"
func FormatTable(headers []string, rows [][]string) string {
	return ""
}

// JoinLines joins lines with newline, no trailing newline when lines is empty.
func JoinLines(lines []string) string {
	return ""
}
