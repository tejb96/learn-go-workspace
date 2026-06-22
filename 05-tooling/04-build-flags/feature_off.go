//go:build !feature

package buildflags

// FeatureFlag reports whether the "feature" build tag is enabled.
func FeatureFlag() string {
	return "off"
}
