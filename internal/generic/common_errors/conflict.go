package common_errors

type ConflictError struct{}

func (e *ConflictError) Error() string {
	return "Failed to create new record ! Duplicate Key"
}
