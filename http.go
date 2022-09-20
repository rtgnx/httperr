package httperr

import "fmt"

type HTTPErr struct {
	Code   int
	Msg    string
	Result interface{}
}

func NewWithData(code int, data interface{}, format string, varg ...interface{}) error {
	err, _ := New(code, format, varg...).(*HTTPErr)
	err.Result = data

	return err
}

func New(code int, format string, varg ...interface{}) error {
	return &HTTPErr{code, fmt.Sprintf(format, varg...), nil}
}

func (h *HTTPErr) Error() string {
	return fmt.Sprintf("HTTP %d: %s", h.Code, h.Msg)
}
