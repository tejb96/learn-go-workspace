package variables

// Swap returns b, a — exchanging two string values without a temporary variable
// is not required; multiple return values are idiomatic in Go.
func Swap(a, b string) (string, string) {
	return "", ""
}

// DescribeType returns a short description of the zero value for the given kind.
// kind must be one of: "string", "int", "bool", "slice".
func DescribeType(kind string) (string, error) {
	return "", nil
}

// DefaultGreeting returns a greeting for name. When name is empty, return
// "Hello, stranger!" — this exercises the string zero value.
func DefaultGreeting(name string) string {
	return ""
}
