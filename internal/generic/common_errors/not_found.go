package common_errors

type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return "Unable to find Record!"
}
