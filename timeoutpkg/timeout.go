package timeoutpkg

import "time"

// TimeoutError type representing a timeout
type TimeoutError struct {
	message string
}

// NewTimeoutError creates a new TimeoutError with a given message
func NewTimeoutError(message string) *TimeoutError {
	return &TimeoutError{message: message}
}

// Error method to implement the error interface
func (e *TimeoutError) Error() string {
	return e.message
}

// Timeout method to signal that this error is due to a timeout
func (e *TimeoutError) Timeout() bool {
	return true
}

// SimulateOperation that may return a timeout error
func SimulateOperation() error {
	// Simulate a delay that could cause a timeout
	time.Sleep(5 * time.Second)
	return NewTimeoutError("Operation timed out")
}
