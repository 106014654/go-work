package errors

import "fmt"

type Error struct {
	ret int32
	data string
	msg string
	cause error
}



func (e *Error) Error() string {
	return fmt.Sprintf("error: ret = %d data = %s message = %s  cause = %v", e.ret, e.data, e.msg,  e.cause)
}

func (e *Error) Unwrap() error { return e.cause }

func (e *Error) WithCause(cause error) *Error {
	err := Clone(e)
	err.cause = cause
	return err
}

func Clone(err *Error) *Error {
	return &Error{
		cause: err.cause,
		ret : err.ret,
		data : err.data,
		msg : err.msg,
	}
}


func New(ret int32, data, message string) *Error {
	return &Error{
		ret : ret,
		data : data,
		msg : message,
	}
}