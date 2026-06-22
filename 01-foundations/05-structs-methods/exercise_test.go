// DO NOT EDIT — implement the solution in solution.go

package structs_test

import (
	"testing"

	"github.com/yourname/go-course/foundations/structs"
)

func TestRectangle_Area(t *testing.T) {
	tests := []struct {
		name           string
		width, height  float64
		want           float64
	}{
		{name: "square", width: 3, height: 3, want: 9},
		{name: "zero dimensions", width: 0, height: 5, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := structs.Rectangle{Width: tt.width, Height: tt.height}
			if got := r.Area(); got != tt.want {
				t.Fatalf("Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_Scale_PointerReceiverMutates(t *testing.T) {
	r := structs.Rectangle{Width: 2, Height: 3}
	r.Scale(2)
	if r.Width != 4 || r.Height != 6 {
		t.Fatalf("after Scale(2): %+v, want Width=4 Height=6", r)
	}
}

func TestRectangle_Scale_ValueReceiverWouldNotMutate(t *testing.T) {
	// Document expected behavior: pointer receiver required.
	r := &structs.Rectangle{Width: 1, Height: 1}
	r.Scale(3)
	if r.Width != 3 {
		t.Fatalf("pointer Scale failed: Width=%v", r.Width)
	}
}

func TestNewRectangle(t *testing.T) {
	tests := []struct {
		name      string
		w, h      float64
		wantErr   bool
		wantW     float64
	}{
		{name: "valid", w: 4, h: 5, wantW: 4},
		{name: "negative width", w: -1, h: 5, wantErr: true},
		{name: "negative height", w: 1, h: -1, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := structs.NewRectangle(tt.w, tt.h)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if r.Width != tt.wantW || r.Height != tt.h {
				t.Fatalf("NewRectangle() = %+v, want w=%v h=%v", r, tt.wantW, tt.h)
			}
		})
	}
}
