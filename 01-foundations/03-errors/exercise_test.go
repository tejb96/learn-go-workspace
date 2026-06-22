// DO NOT EDIT — implement the solution in solution.go

package errorslesson_test

import (
	"errors"
	"fmt"
	"testing"

	errs "github.com/yourname/go-course/foundations/errors"
)

func TestFind_HappyPath(t *testing.T) {
	store := map[string]string{"u1": "Ada"}
	got, err := errs.Find(store, "u1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "Ada" {
		t.Fatalf("Find() = %q, want Ada", got)
	}
}

func TestFind_NotFound(t *testing.T) {
	_, err := errs.Find(map[string]string{}, "missing")
	if !errors.Is(err, errs.ErrNotFound) {
		t.Fatalf("error = %v, want ErrNotFound", err)
	}
}

func TestFind_EmptyIDValidation(t *testing.T) {
	_, err := errs.Find(map[string]string{"u1": "Ada"}, "")
	var ve *errs.ValidationError
	if !errors.As(err, &ve) {
		t.Fatalf("error = %T %v, want *ValidationError", err, err)
	}
	if ve.Field != "id" || ve.Code != "required" {
		t.Fatalf("ValidationError = %+v, want Field=id Code=required", ve)
	}
}

func TestWrap_Unwrap(t *testing.T) {
	inner := errs.ErrNotFound
	wrapped := errs.Wrap(inner, "lookup user")
	if !errors.Is(wrapped, errs.ErrNotFound) {
		t.Fatalf("errors.Is(wrapped, ErrNotFound) = false, want true")
	}
	if got := wrapped.Error(); got != "lookup user: not found" {
		t.Fatalf("wrapped.Error() = %q, want %q", got, "lookup user: not found")
	}
}

func TestValidationError_As(t *testing.T) {
	err := fmt.Errorf("outer: %w", &errs.ValidationError{Field: "email", Code: "invalid"})
	var ve *errs.ValidationError
	if !errors.As(err, &ve) {
		t.Fatal("errors.As failed")
	}
	if ve.Field != "email" {
		t.Fatalf("Field = %q, want email", ve.Field)
	}
}
