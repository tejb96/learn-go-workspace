package variables

import "fmt"

// Swap returns b, a — exchanging two string values without a temporary variable
// is not required; multiple return values are idiomatic in Go.
func Swap(a, b string) (string, string) {
	return b, a
}

// DescribeType returns a short description of the zero value for the given kind.
// kind must be one of: "string", "int", "bool", "slice".
func DescribeType(kind string) (string, error) {
	switch kind {
	case "string":
		return "zero value is \"\"", nil
	case "int":
		return "zero value is 0", nil
	case "bool":
		return "zero value is false", nil
	case "slice":
		return "zero value is nil", nil
	default:
		return "", fmt.Errorf("Unknown or unsupported kind: %v", kind)
	}
}

// DefaultGreeting returns a greeting for name. When name is empty, return
// "Hello, stranger!" — this exercises the string zero value.
func DefaultGreeting(name string) string {
	switch name{
		case "":
			return "Hello, stranger!"
		default:
			return fmt.Sprintf("Hello, %v!",name)
	}
}
