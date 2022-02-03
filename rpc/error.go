package rpc

import (
	"errors"
	"fmt"
)

type Error struct {
	error
	Code int
}

func NewError(code int, msg ...string) Error {
	message := ""
	if len(msg) > 0 {
		message = msg[0]
	}
	return Error{
		error: errors.New(fmt.Sprintf("rpc error: %d. %s", code, message)),
		Code:  code,
	}
}

const (
	ErrorEmpty                   = 0
	ErrorFunctionNotCallable     = -100
	ErrorFunctionNotFound        = -101
	ErrorIncompatibleRPCVersion  = -102
	ErrorInvalidArguments        = -200
	ErrorEncode                  = -300
	ErrorDecode                  = -301
	ErrorCouldNotSendCall        = -400
	ErrorNotSuccessfulStatusCode = -401
	ErrorServer                  = -500
	ErrorAuth                    = -600
	ErrorPreCallChecks           = -700
)
