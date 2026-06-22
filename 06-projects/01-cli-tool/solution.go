package clitool

import (
	"io"
)

// LogEntry is one parsed log line.
type LogEntry struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

// ParseLine parses a line in the form "LEVEL message text".
// LEVEL is the first field; the rest is the message.
func ParseLine(line string) (LogEntry, error) {
	return LogEntry{}, nil
}

// FilterByLevel returns entries matching level (case-insensitive). Empty level returns all.
func FilterByLevel(entries []LogEntry, level string) []LogEntry {
	return nil
}

// Run executes the log parser CLI.
// Flags: -level LEVEL, -format plain|json, -file PATH (optional; default stdin).
func Run(args []string, stdin io.Reader, stdout io.Writer) error {
	return nil
}
