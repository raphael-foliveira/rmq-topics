package res

type HttpError struct {
	Status  int
	Message string
}

func (e *HttpError) Error() string {
	return e.Message
}
