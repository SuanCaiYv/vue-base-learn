package resp

import "time"

var notFound = &Result{
	Code: 404,
	Msg:  "Not Fount",
	Data: struct {
	}{},
	Timestamp: time.Now(),
}

var missToken = &Result{
	Code: 451,
	Msg:  "miss token",
	Data: struct {
	}{},
	Timestamp: time.Now(),
}

var authFailed = &Result{
	Code: 452,
	Msg:  "auth failed",
	Data: struct {
	}{},
	Timestamp: time.Now(),
}

type Result struct {
	Code      int
	Msg       string
	Data      interface{}
	Timestamp time.Time
}

func NewOk(data interface{}) *Result {
	return &Result{
		Code:      200,
		Msg:       "",
		Data:      data,
		Timestamp: time.Now(),
	}
}

func NewNotFound() *Result {
	return notFound
}

func NewIntervalError(msg string) *Result {
	return &Result{
		Code: 500,
		Msg:  msg,
		Data: struct {
		}{},
		Timestamp: time.Now(),
	}
}

func NewMissToken() *Result {
	return missToken
}

func NewAuthFailed() *Result {
	return authFailed
}

func NewBadRequest(msg string) *Result {
	return &Result{
		Code: 400,
		Msg:  msg,
		Data: struct {
		}{},
		Timestamp: time.Now(),
	}
}
