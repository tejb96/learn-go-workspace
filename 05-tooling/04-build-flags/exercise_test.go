// DO NOT EDIT — implement the solution in solution.go and feature_*.go

package buildflags_test

import (
	"testing"

	"github.com/yourname/go-course/tooling/buildflags"
)

func TestFeatureFlag_Default(t *testing.T) {
	if got := buildflags.FeatureFlag(); got != "off" {
		t.Fatalf("FeatureFlag() = %q, want off (default build)", got)
	}
}

func TestDebugMode(t *testing.T) {
	if got := buildflags.DebugMode(); got != "default" {
		t.Fatalf("DebugMode() = %q, want default", got)
	}
}
