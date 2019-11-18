package basic

import "fmt"

type HttpRequestError string

func (h HttpRequestError) Error() string {
	return fmt.Sprintf("http request error: %s", string(h))
}
