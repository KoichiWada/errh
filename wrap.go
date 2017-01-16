package errh

import (
	"bytes"
	"fmt"
)

// Wrap returns an error that formats as err with the point Wrap is called, and the supplied message.
// Wrap also records the cause of err.
// If err is nil, Wrap returns nil.
func Wrap(err error, a ...interface{}) error {
	if err == nil {
		return nil
	}

	buf := new(bytes.Buffer)
	if len(a) > 0 {
		format := fmt.Sprintf("%v; ", a[0])
		fmt.Fprintf(buf, format, a[1:]...)
	}
	buf.WriteString(err.Error())
	addFileLine(buf)

	return &wrapper{
		cause: Cause(err),
		msg:   buf.String(),
	}
}

type wrapper struct {
	cause error
	msg   string
}

// Error
func (w *wrapper) Error() string {
	return w.msg
}

// Cause
func (w *wrapper) Cause() error {
	return w.cause
}

type causer interface {
	Cause() error
}

// Cause returns the underlying cause of the error, if possible.
func Cause(err error) error {
	if err == nil {
		return nil
	}

	c, ok := err.(causer)
	if !ok {
		return err
	}

	return c.Cause()
}
