package v2

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidRequestObj = errors.New("failed to build request")
	ErrNotFound          = errors.New("object not found")
)

type (
	BadResponseError struct {
		ErrorMsg string `json:"error,omitempty"` //nolint: tagliatelle
		Code     int    `json:"code"`
	}
)

func (e BadResponseError) Error() string {
	return fmt.Sprintf("error response: %v", e.ErrorMsg)
}
