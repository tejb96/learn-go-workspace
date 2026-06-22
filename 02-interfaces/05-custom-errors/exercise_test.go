// DO NOT EDIT — implement the solution in solution.go

package customerrors_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/yourname/go-course/interfaces/customerrors"
)

func TestOpError_ErrorAndUnwrap(t *testing.T) {
	inner := errors.New("disk full")
	err := customerrors.NewOpError("write", inner)
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got := err.Error(); got != "write: disk full" {
		t.Fatalf("Error() = %q, want %q", got, "write: disk full")
	}
	if !errors.Is(err, inner) {
		t.Fatal("errors.Is should find inner")
	}
	var oe *customerrors.OpError
	if !errors.As(err, &oe) {
		t.Fatal("errors.As should find OpError")
	}
	if oe.Op != "write" {
		t.Fatalf("Op = %q, want write", oe.Op)
	}
}

func TestNewOpError_NilErr(t *testing.T) {
	if got := customerrors.NewOpError("read", nil); got != nil {
		t.Fatalf("NewOpError with nil err = %v, want nil", got)
	}
}

func TestOpFromError(t *testing.T) {
	err := fmt.Errorf("outer: %w", customerrors.NewOpError("delete", errors.New("x")))
	if got := customerrors.OpFromError(err); got != "delete" {
		t.Fatalf("OpFromError = %q, want delete", got)
	}
	if got := customerrors.OpFromError(errors.New("plain")); got != "" {
		t.Fatalf("OpFromError plain = %q, want empty", got)
	}
}

func TestMarkRetryable(t *testing.T) {
	inner := errors.New("timeout")
	err := customerrors.MarkRetryable(inner, "network blip")
	if !errors.Is(err, customerrors.ErrRetryable) {
		t.Fatal("expected ErrRetryable in chain")
	}
	var re *customerrors.RetryableError
	if !errors.As(err, &re) {
		t.Fatal("expected RetryableError")
	}
	if re.Reason != "network blip" {
		t.Fatalf("Reason = %q, want network blip", re.Reason)
	}
}
