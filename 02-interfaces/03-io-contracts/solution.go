package iocontracts

import "io"

// CopyN copies up to n bytes from src to dst using io.CopyN semantics.
// Returns bytes written and any error from the underlying readers/writers.
func CopyN(dst io.Writer, src io.Reader, n int64) (written int64, err error) {
	return 0, nil
}

// ReadAll reads everything from r until EOF.
func ReadAll(r io.Reader) ([]byte, error) {
	return nil, nil
}

// WriteString writes s to w and returns bytes written.
func WriteString(w io.Writer, s string) (int, error) {
	return 0, nil
}
