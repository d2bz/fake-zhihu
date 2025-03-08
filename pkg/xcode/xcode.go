package xcode

import "strconv"

type Xcode interface {
	Error() string
	Code() int
	Message() string
	Detail() []interface{}
}

type Err struct {
	code int
	msg  string
}

func (e Err) Error() string {
	if len(e.msg) > 0 {
		return e.msg
	}

	return strconv.Itoa(e.code)
}

func (e Err) Code() int {
	return e.code
}

func (e Err) Message() string {
	return e.Error()
}

func (e Err) Detail() []interface{} {
	return nil
}

func NewErr(code int, msg string) Err {
	return Err{
		code: code,
		msg:  msg,
	}
}

func Define(code int, msg string) Err {
	return Err{
		code: code,
		msg:  msg,
	}
}
