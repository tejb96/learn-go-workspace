// DO NOT EDIT — implement the solution in solution.go

package iocontracts_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/yourname/go-course/interfaces/iocontracts"
)

func TestCopyN_BufferAndStringsReader(t *testing.T) {
	src := strings.NewReader("hello world")
	dst := &bytes.Buffer{}

	written, err := iocontracts.CopyN(dst, src, 5)
	if err != nil {
		t.Fatalf("CopyN error: %v", err)
	}
	if written != 5 {
		t.Fatalf("written = %d, want 5", written)
	}
	if dst.String() != "hello" {
		t.Fatalf("dst = %q, want hello", dst.String())
	}
}

func TestCopyN_FileToBuffer(t *testing.T) {
	dir := t.TempDir()
	path := dir + "/data.txt"
	if err := os.WriteFile(path, []byte("file content"), 0o644); err != nil {
		t.Fatal(err)
	}

	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	dst := &bytes.Buffer{}
	written, err := iocontracts.CopyN(dst, f, 4)
	if err != nil {
		t.Fatalf("CopyN error: %v", err)
	}
	if written != 4 || dst.String() != "file" {
		t.Fatalf("got written=%d dst=%q", written, dst.String())
	}
}

func TestCopyN_ExceedsSource(t *testing.T) {
	src := strings.NewReader("ab")
	dst := &bytes.Buffer{}
	written, err := iocontracts.CopyN(dst, src, 100)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if written != 2 || dst.String() != "ab" {
		t.Fatalf("written=%d dst=%q", written, dst.String())
	}
}

func TestReadAll(t *testing.T) {
	src := strings.NewReader("stream")
	got, err := iocontracts.ReadAll(src)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "stream" {
		t.Fatalf("ReadAll = %q, want stream", got)
	}
}

func TestWriteString(t *testing.T) {
	var buf bytes.Buffer
	n, err := iocontracts.WriteString(&buf, "go")
	if err != nil {
		t.Fatal(err)
	}
	if n != 2 || buf.String() != "go" {
		t.Fatalf("n=%d buf=%q", n, buf.String())
	}
}
