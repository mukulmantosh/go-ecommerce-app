package common_errors

type ViolationError struct{}

func (e *ViolationError) Error() string {
	return "Key Violation Error"
}
