package customerrors

import "errors"

// OpError records which operation failed and wraps an underlying error.
type OpError struct {
	Op  string
	Err error
}

func (e *OpError) Error() string {
	return ""
}

func (e *OpError) Unwrap() error {
	return nil
}

// NewOpError returns an OpError when err is non-nil. Returns nil when err is nil.
func NewOpError(op string, err error) error {
	return err
}

// OpFromError returns the Op field if any OpError appears in err's chain.
func OpFromError(err error) string {
	return ""
}

// IsRetryable reports whether err wraps errors marked with ErrRetryable.
var ErrRetryable = errors.New("retryable")

// RetryableError wraps ErrRetryable for use with errors.Is.
type RetryableError struct {
	Reason string
}

func (e *RetryableError) Error() string {
	return ""
}

func (e *RetryableError) Unwrap() error {
	return nil
}

// MarkRetryable wraps err as retryable with a reason.
func MarkRetryable(err error, reason string) error {
	return err
}
