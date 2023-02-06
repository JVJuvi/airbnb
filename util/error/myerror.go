package myerror

import "github.com/pkg/errors"

type MyError struct {
	Raw       error
	ErrorCode string
	HTTPCode  int
	Message   string
}

func (m MyError) Error() string {
	if m.Raw != nil {
		return errors.Wrap(m.Raw, m.Message).Error()
	}

	return m.Message
}
