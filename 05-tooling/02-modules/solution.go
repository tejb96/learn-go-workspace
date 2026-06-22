package moduleslesson

// ModulePath parses the module path from go.mod file contents.
// Returns the path from the first "module" directive line.
func ModulePath(modContents string) (string, error) {
	return "", nil
}

// HasRequire reports whether modContents contains a require line for modulePath.
func HasRequire(modContents, modulePath string) bool {
	return false
}

// WorkUsePaths parses go.work contents and returns all paths from "use" blocks.
func WorkUsePaths(workContents string) ([]string, error) {
	return nil, nil
}

// ChildGreet calls the local child module dependency.
func ChildGreet(name string) string {
	return ""
}
