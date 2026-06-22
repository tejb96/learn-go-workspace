package workerpool

// FileResult holds line count for one file.
type FileResult struct {
	Path  string `json:"path"`
	Lines int    `json:"lines"`
	Error string `json:"error,omitempty"`
}

// ProcessDir reads all regular files in dir (non-recursive) using workers goroutines.
// Returns one FileResult per file, sorted by Path ascending.
func ProcessDir(dir string, workers int) ([]FileResult, error) {
	return nil, nil
}
