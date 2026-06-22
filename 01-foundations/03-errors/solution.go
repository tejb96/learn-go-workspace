package errorslesson

import "errors"

// ErrNotFound is returned when a user ID does not exist in the store.
var ErrNotFound = errors.New("not found")

// ValidationError indicates invalid input. Code is a machine-readable tag.
type ValidationError struct {
	Field string
	Code  string
}

func (e *ValidationError) Error() string {
	return "validation failed"
}

// Wrap wraps err with additional context using fmt.Errorf and %w.
func Wrap(err error, msg string) error {
	return err
}

// Find looks up id in store. Returns ErrNotFound when missing.
// When id is empty, returns a ValidationError with Field "id" and Code "required".
func Find(store map[string]string, id string) (string, error) {
	return "", nil
}
