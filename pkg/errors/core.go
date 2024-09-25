package errors

type PanicError struct {
	error
	Message string
}

func (panicError PanicError) Error() string {
	if panicError.Message == "" {
		return "Panic error"
	}
	return panicError.Message
}

func NewPanicError(message string) *PanicError {
	return &PanicError{Message: message}
}

type ValidatorError struct {
	error
	Message string
}

func (validatorError ValidatorError) Error() string {
	if validatorError.Message == "" {
		return "Validator error"
	}
	return validatorError.Message
}

func NewValidatorError(message string) *ValidatorError {
	return &ValidatorError{Message: message}
}
