package httpErrors

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type ApiErr interface {
	Status() int
	Error() string
}
type ApiError struct {
	ErrStatus int    `json:"code,omitempty"`
	ErrError  string `json:"code,omitempty"`
}

var (
	BadRequest          = errors.New("Bad Request")
	ContentType         = errors.New("Content type must be `application/json`")
	CannotMarshal       = errors.New("Could not be marshalled")
	NotFound            = errors.New("Not Found")
	BadQueryParams      = errors.New("Bad Query Params")
	InternalServerError = errors.New("Internal Server Error")
)

func (a *ApiError) Status() int {
	return a.ErrStatus
}

func (a *ApiError) Error() string {
	return fmt.Sprintf("status: %d - errors: %s", a.ErrStatus, a.ErrError)
}

func NewApiError(code int, err string) ApiError {
	return ApiError{
		ErrStatus: code,
		ErrError:  err,
	}
}

func ParseErrors(err error) ApiError {
	switch {
	case strings.Contains(err.Error(), "json: unsupported"):
		return NewApiError(http.StatusBadRequest, CannotMarshal.Error())
	default:
		return NewApiError(http.StatusBadRequest, NotFound.Error())
	}
}
