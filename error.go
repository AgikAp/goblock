package goblock

import "fmt"

type HttpError struct {
	StatusCode int         `json:"statusCode"`
	Message    interface{} `json:"message"`
	Previous   error       `json:"-"`
}

func NewHttpError(statusCode int, message interface{}, previous error) *HttpError {
	return &HttpError{statusCode, message, previous}
}

func (e HttpError) Error() string {
	if e.Previous != nil {
		return fmt.Sprintf("statusCode=%d, message=%v, previous=%v", e.StatusCode, e.Message, e.Previous)
	}

	return fmt.Sprintf("statusCode=%d, message=%v", e.StatusCode, e.Message)
}
