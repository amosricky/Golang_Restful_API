package util

import (
	"fmt"
	"net/http"
)

type ExtendError interface {
	GetCode() int
	GetMessage() string
	Error() string
}

type BaseError struct {
	Code    int    `json:"error_code"`
	Message string `json:"error_message"`
}

func (b *BaseError) Error() string {
	return fmt.Sprintf("Code:%d: Message:%s", b.Code, b.Message)
}

func (b *BaseError) GetCode() int {
	return b.Code
}

func (b *BaseError) GetMessage() string {
	return b.Message
}

func NewBaseError(code int, message string) ExtendError {

	if len(message) == 0 {
		if 10000 <= code {
			message = GetMsg(code)
		} else {
			message = http.StatusText(code)
		}
	}

	return &BaseError{Code: code, Message: message }
}
