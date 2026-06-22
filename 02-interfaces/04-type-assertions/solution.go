package typeassertions

import "errors"

var ErrUnsupportedType = errors.New("unsupported type")

// ParseAny converts supported dynamic types to a canonical string form.
// Supported: string (as-is), int (decimal), bool ("true"/"false").
// Prefer concrete types or generics in real code — any is a last resort.
func ParseAny(v any) (string, error) {
	return "", nil
}

// AsInt extracts an int from v when possible (int or int64 that fits int).
func AsInt(v any) (int, bool) {
	return 0, false
}
