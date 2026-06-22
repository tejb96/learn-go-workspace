// DO NOT EDIT — implement the solution in solution.go

package typeassertions_test

import (
	"errors"
	"testing"

	"github.com/yourname/go-course/interfaces/typeassertions"
)

func TestParseAny(t *testing.T) {
	tests := []struct {
		name    string
		in      any
		want    string
		wantErr error
	}{
		{name: "string", in: "hello", want: "hello"},
		{name: "int", in: 42, want: "42"},
		{name: "bool true", in: true, want: "true"},
		{name: "bool false", in: false, want: "false"},
		{name: "unsupported", in: 3.14, wantErr: typeassertions.ErrUnsupportedType},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := typeassertions.ParseAny(tt.in)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("error = %v, want %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("ParseAny() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestAsInt(t *testing.T) {
	tests := []struct {
		name   string
		in     any
		want   int
		wantOK bool
	}{
		{name: "int", in: 7, want: 7, wantOK: true},
		{name: "int64", in: int64(9), want: 9, wantOK: true},
		{name: "string fails", in: "7", wantOK: false},
		{name: "float fails", in: 1.5, wantOK: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := typeassertions.AsInt(tt.in)
			if ok != tt.wantOK {
				t.Fatalf("ok = %v, want %v", ok, tt.wantOK)
			}
			if ok && got != tt.want {
				t.Fatalf("AsInt() = %d, want %d", got, tt.want)
			}
		})
	}
}
