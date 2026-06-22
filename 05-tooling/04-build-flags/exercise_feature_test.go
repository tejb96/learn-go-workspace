//go:build feature

// DO NOT EDIT — run with: go test -tags=feature -run TestFeatureFlag_Enabled

package buildflags_test

import (
	"testing"

	"github.com/yourname/go-course/tooling/buildflags"
)

func TestFeatureFlag_Enabled(t *testing.T) {
	if got := buildflags.FeatureFlag(); got != "on" {
		t.Fatalf("FeatureFlag() = %q, want on (-tags=feature)", got)
	}
}
