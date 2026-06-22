package fileio

// ReadLines reads path and returns lines without trailing newline characters.
// Empty file returns empty slice, not nil.
func ReadLines(path string) ([]string, error) {
	return nil, nil
}

// WriteAtomic writes data to path atomically via a temp file and rename.
// Creates parent directories if needed. On failure, leaves original path unchanged.
func WriteAtomic(path string, data []byte) error {
	return nil
}

// FileExists reports whether path exists (any type, including directories).
func FileExists(path string) bool {
	return false
}
