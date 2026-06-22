// DO NOT EDIT — implement the solution in solution.go

package moduleslesson_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/yourname/go-course/tooling/moduleslesson"
)

func TestModulePath(t *testing.T) {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		t.Fatal(err)
	}
	got, err := moduleslesson.ModulePath(string(data))
	if err != nil {
		t.Fatal(err)
	}
	want := "github.com/yourname/go-course/tooling/moduleslesson"
	if got != want {
		t.Fatalf("ModulePath() = %q, want %q", got, want)
	}
}

func TestHasRequire(t *testing.T) {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		t.Fatal(err)
	}
	if !moduleslesson.HasRequire(string(data), "github.com/yourname/go-course/tooling/moduleschild") {
		t.Fatal("expected require for moduleschild")
	}
	if moduleslesson.HasRequire(string(data), "github.com/example/missing") {
		t.Fatal("unexpected require match")
	}
}

func TestWorkUsePaths(t *testing.T) {
	data, err := os.ReadFile("go.work")
	if err != nil {
		t.Fatal(err)
	}
	got, err := moduleslesson.WorkUsePaths(string(data))
	if err != nil {
		t.Fatal(err)
	}
	want := []string{".", "./examples/child"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("WorkUsePaths() = %v, want %v", got, want)
	}
}

func TestChildModulePath(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("examples", "child", "go.mod"))
	if err != nil {
		t.Fatal(err)
	}
	got, err := moduleslesson.ModulePath(string(data))
	if err != nil {
		t.Fatal(err)
	}
	want := "github.com/yourname/go-course/tooling/moduleschild"
	if got != want {
		t.Fatalf("child ModulePath = %q, want %q", got, want)
	}
}

func TestChildGreet(t *testing.T) {
	got := moduleslesson.ChildGreet("Go")
	if got != "hello Go" {
		t.Fatalf("ChildGreet() = %q, want %q", got, "hello Go")
	}
}
