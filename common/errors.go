package common

import "fmt"

type restyHttpError struct {
	msg string
}

func NewRestyHttpError(response []byte, statusCode int, err error) *restyHttpError {
	var r restyHttpError
	r.msg = fmt.Sprintf("an error occured while sending http request: status code: %v, error: %s ", statusCode, err.Error())
	if len(response) != 0 {
		r.msg += fmt.Sprintf("http response: %s", string(response))
	}

	return &r
}

func (r *restyHttpError) Error() string {
	return r.msg
}
