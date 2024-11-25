package errors

import (
	"fmt"

	"github.com/zawa-t/pr/reporter/src/env/lang"
)

type errorCode int

const (
	_ errorCode = iota
	NotFound
	InvalidParams
)

var errorMessages = map[errorCode]map[int]string{
	NotFound: {
		lang.JPN: "見つかりませんでした。",
		lang.ENG: "Not found.",
	},
	InvalidParams: {
		lang.JPN: "指定された内容に誤りがあります。",
		lang.ENG: "Invalid parameter.",
	},
}

type AppError struct {
	Code errorCode
	Err  error
}

func NewAppError(code errorCode, err error) *AppError {
	return &AppError{code, err}
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s (ErrorCode: %d)", e.Err.Error(), e.Code)
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) Is(err error) bool {
	appErr, ok := err.(*AppError)
	return ok && e.Code == appErr.Code
}

func (e *AppError) As(target any) bool {
	if ptr, ok := target.(**AppError); ok {
		*ptr = e
		return true
	}
	return false
}

func (e *AppError) Message() string {
	if msgs, ok := errorMessages[e.Code]; ok {
		if msg, ok := msgs[lang.Language()]; ok {
			return msg
		}
	}
	return ""
}
