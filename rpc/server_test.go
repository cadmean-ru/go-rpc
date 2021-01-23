package rpc

import (
	"net/http"
	"testing"
)

func TestHandle(t *testing.T) {
	Handle("sum", func(args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, NewError(1, "Not enough arguments")
		}

		a := args[0].(float64)
		b := args[1].(float64)

		return a+b, nil
	})

	_ = http.ListenAndServe(":69", nil)
}