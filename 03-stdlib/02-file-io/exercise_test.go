// DO NOT EDIT — implement the solution in solution.go

package fileio_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/yourname/go-course/stdlib/fileio"
)

func TestReadLines(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "lines.txt")
	content := "first\nsecond\nthird\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	got, err := fileio.ReadLines(path)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"first", "second", "third"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ReadLines() = %v, want %v", got, want)
	}
}

func TestReadLines_EmptyFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "empty.txt")
	if err := os.WriteFile(path, nil, 0o644); err != nil {
		t.Fatal(err)
	}
	got, err := fileio.ReadLines(path)
	if err != nil {
		t.Fatal(err)
	}
	if got == nil || len(got) != 0 {
		t.Fatalf("ReadLines empty = %v, want empty slice", got)
	}
}

func TestReadLines_MissingFile(t *testing.T) {
	_, err := fileio.ReadLines(filepath.Join(t.TempDir(), "nope.txt"))
	if err == nil {
		t.Fatal("expected error for missing file")
	}
}

func TestWriteAtomic(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "sub", "out.txt")

	if err := fileio.WriteAtomic(path, []byte("hello")); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "hello" {
		t.Fatalf("file = %q, want hello", data)
	}
}

func TestWriteAtomic_DoesNotCorruptOnFailure(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "keep.txt")
	if err := os.WriteFile(path, []byte("original"), 0o644); err != nil {
		t.Fatal(err)
	}

	// Writing to path inside a non-existent root we cannot create should fail on some systems;
	// use invalid path under a file (not a directory) to force failure.
	blocker := filepath.Join(dir, "blocker")
	if err := os.WriteFile(blocker, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	badPath := filepath.Join(blocker, "out.txt")

	err := fileio.WriteAtomic(badPath, []byte("bad"))
	if err == nil {
		t.Fatal("expected error writing under file path")
	}
	data, _ := os.ReadFile(path)
	if string(data) != "original" {
		t.Fatalf("original corrupted = %q", data)
	}
}

func TestFileExists(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "present.txt")
	if err := os.WriteFile(path, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if !fileio.FileExists(path) {
		t.Fatal("expected file to exist")
	}
	if fileio.FileExists(filepath.Join(dir, "missing.txt")) {
		t.Fatal("expected missing file to return false")
	}
}
