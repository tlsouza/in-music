package errors

func UnauthorizedError(err error) *HttpError {
	return NewHttpError(err, 401)
}
