// DO NOT EDIT — implement the solution in solution.go

package variables_test

import (
	"testing"

	"github.com/yourname/go-course/foundations/variables"
)

func TestSwap(t *testing.T) {
	tests := []struct {
		name string
		a, b string
		wantA, wantB string
	}{
		{name: "two names", a: "hello", b: "world", wantA: "world", wantB: "hello"},
		{name: "empty first", a: "", b: "go", wantA: "go", wantB: ""},
		{name: "both empty", a: "", b: "", wantA: "", wantB: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA, gotB := variables.Swap(tt.a, tt.b)
			if gotA != tt.wantA || gotB != tt.wantB {
				t.Fatalf("Swap(%q, %q) = (%q, %q), want (%q, %q)", tt.a, tt.b, gotA, gotB, tt.wantA, tt.wantB)
			}
		})
	}
}

func TestDescribeType(t *testing.T) {
	tests := []struct {
		name    string
		kind    string
		want    string
		wantErr bool
	}{
		{name: "string zero", kind: "string", want: `zero value is ""`},
		{name: "int zero", kind: "int", want: "zero value is 0"},
		{name: "bool zero", kind: "bool", want: "zero value is false"},
		{name: "slice zero", kind: "slice", want: "zero value is nil"},
		{name: "unknown kind", kind: "chan", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := variables.DescribeType(tt.kind)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("DescribeType(%q) = %q, want %q", tt.kind, got, tt.want)
			}
		})
	}
}

func TestDefaultGreeting_EmptyNameUsesZeroValue(t *testing.T) {
	got := variables.DefaultGreeting("")
	if got != "Hello, stranger!" {
		t.Fatalf("DefaultGreeting(\"\") = %q, want %q", got, "Hello, stranger!")
	}
}

func TestDefaultGreeting_WithName(t *testing.T) {
	got := variables.DefaultGreeting("Ada")
	if got != "Hello, Ada!" {
		t.Fatalf("DefaultGreeting(%q) = %q, want %q", "Ada", got, "Hello, Ada!")
	}
}
