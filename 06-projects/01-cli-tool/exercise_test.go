// DO NOT EDIT — implement the solution in solution.go

package clitool_test

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/yourname/go-course/projects/clitool"
)

func TestParseLine(t *testing.T) {
	tests := []struct {
		name    string
		line    string
		want    clitool.LogEntry
		wantErr bool
	}{
		{name: "info", line: "INFO server started", want: clitool.LogEntry{Level: "INFO", Message: "server started"}},
		{name: "error with spaces", line: "ERROR disk full on /data", want: clitool.LogEntry{Level: "ERROR", Message: "disk full on /data"}},
		{name: "invalid", line: "   ", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := clitool.ParseLine(tt.line)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Fatalf("ParseLine() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestFilterByLevel(t *testing.T) {
	entries := []clitool.LogEntry{
		{Level: "INFO", Message: "a"},
		{Level: "ERROR", Message: "b"},
		{Level: "info", Message: "c"},
	}
	got := clitool.FilterByLevel(entries, "info")
	if len(got) != 2 {
		t.Fatalf("FilterByLevel() len = %d, want 2", len(got))
	}
}

func TestRun_StdinPlain(t *testing.T) {
	in := strings.NewReader("INFO hello\nERROR oops\n")
	var out bytes.Buffer
	err := clitool.Run([]string{"-level=ERROR", "-format=plain"}, in, &out)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), "ERROR oops") {
		t.Fatalf("output = %q", out.String())
	}
	if strings.Contains(out.String(), "hello") {
		t.Fatal("INFO line should be filtered out")
	}
}

func TestRun_FileJSON(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "app.log")
	content := "WARN low memory\nINFO ok\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	var out bytes.Buffer
	err := clitool.Run([]string{"-file=" + path, "-level=WARN", "-format=json"}, nil, &out)
	if err != nil {
		t.Fatal(err)
	}

	var entries []clitool.LogEntry
	if err := json.Unmarshal(out.Bytes(), &entries); err != nil {
		t.Fatalf("invalid JSON: %v body=%q", err, out.String())
	}
	if len(entries) != 1 || entries[0].Level != "WARN" {
		t.Fatalf("entries = %+v", entries)
	}
}

func TestRun_MissingFile(t *testing.T) {
	var out bytes.Buffer
	err := clitool.Run([]string{"-file=/no/such/file.log"}, nil, &out)
	if err == nil {
		t.Fatal("expected error for missing file")
	}
}
