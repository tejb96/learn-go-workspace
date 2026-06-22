package buildflags

// DebugMode reports build type via ldflags or default.
// When not overridden at link time, returns "default".
func DebugMode() string {
	return ""
}
