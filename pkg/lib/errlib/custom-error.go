package errlib

type CustomError struct {
	message string
}

func NewCustomError(msg string) *CustomError {
	return &CustomError{message: msg}
}

func (e *CustomError) Error() string {
	return e.message
}

func IsCustom(err error) bool {
	if _, ok := err.(*CustomError); ok {
		return true
	}

	return false
}
