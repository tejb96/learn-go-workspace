package composition

// Reader can produce bytes.
type Reader interface {
	Read() []byte
}

// Writer can consume bytes.
type Writer interface {
	Write(p []byte) int
}

// ReadWriter combines Reader and Writer via interface embedding.
type ReadWriter interface {
	Reader
	Writer
}

// Describe reports which capabilities v has: "reader", "writer", "readwriter", or "none".
// A type may satisfy multiple interfaces; prefer the most capable label ("readwriter" beats "reader").
func Describe(v any) string {
	return ""
}

// Duplicate reads from r and writes the same bytes to each writer. Returns total bytes written per writer.
func Duplicate(r Reader, writers ...Writer) []int {
	return nil
}
