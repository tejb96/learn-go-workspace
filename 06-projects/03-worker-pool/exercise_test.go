// DO NOT EDIT — implement the solution in solution.go

package workerpool_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/yourname/go-course/projects/workerpool"
)

func TestProcessDir(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "b.txt"), "line1\nline2\n")
	writeFile(t, filepath.Join(dir, "a.txt"), "only\n")
	writeFile(t, filepath.Join(dir, "empty.txt"), "")

	got, err := workerpool.ProcessDir(dir, 2)
	if err != nil {
		t.Fatal(err)
	}

	want := []workerpool.FileResult{
		{Path: "a.txt", Lines: 1},
		{Path: "b.txt", Lines: 2},
		{Path: "empty.txt", Lines: 0},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ProcessDir() = %+v, want %+v", got, want)
	}
}

func TestProcessDir_SkipsSubdirectories(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "ok.txt"), "x\n")
	if err := os.Mkdir(filepath.Join(dir, "sub"), 0o755); err != nil {
		t.Fatal(err)
	}

	got, err := workerpool.ProcessDir(dir, 1)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 || got[0].Path != "ok.txt" {
		t.Fatalf("got %+v, want only ok.txt", got)
	}
}

func TestProcessDir_MissingDir(t *testing.T) {
	_, err := workerpool.ProcessDir(filepath.Join(t.TempDir(), "missing"), 2)
	if err == nil {
		t.Fatal("expected error for missing directory")
	}
}

func TestProcessDir_InvalidWorkers(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "a.txt"), "x\n")
	_, err := workerpool.ProcessDir(dir, 0)
	if err == nil {
		t.Fatal("expected error for workers < 1")
	}
}

func writeFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}
