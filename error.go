package doku

func (e ErrorResponse) Error() string {
	return e.ErrorDetail.Message
}
