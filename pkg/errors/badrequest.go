package errors

func BadRequestError(err error) *HttpError {
	return NewHttpError(err, 400)
}
