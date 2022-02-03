package server

import (
	"encoding/json"
	"fmt"
	"github.com/cadmean-ru/go-rpc/rpc"
	"io/ioutil"
	"net/http"
)

const Version = "2.1"

type HandlerFunc func(args ...interface{}) (interface{}, error)

func MakeHandler(delegate HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS")
		writer.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
		writer.Header().Add("Content-Type", "application/json")

		v := request.Header.Get("Cadmean-RPC-Version")
		if v != Version {
			writeError(writer, rpc.ErrorIncompatibleRPCVersion)
			return
		}

		defer request.Body.Close()
		data, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writeError(writer, rpc.ErrorDecode)
			return
		}

		call := rpc.FunctionCall{}
		err = json.Unmarshal(data, &call)
		if err != nil {
			writeError(writer, rpc.ErrorDecode)
			return
		}

		defer func() {
			recoveredError := recover()
			if recoveredError != nil {
				writeError(writer, rpc.ErrorServer)
			}
		}()

		result, customError := delegate(call.Arguments...)
		var errorCode int
		if customError != nil {
			if rpcError, ok := customError.(rpc.Error); ok {
				errorCode = rpcError.Code
			} else {
				errorCode = rpc.ErrorServer
			}
		}

		output := rpc.FunctionOutput{
			Error:  errorCode,
			Result: result,
			Meta:   nil,
		}

		responseData, err := json.Marshal(output)
		if err != nil {
			writeError(writer, rpc.ErrorEncode)
			return
		}

		_, _ = writer.Write(responseData)
	}
}

func Handle(functionName string, handler HandlerFunc) {
	http.HandleFunc(fmt.Sprintf("/api/rpc/%s", functionName), MakeHandler(handler))
}

func writeError(writer http.ResponseWriter, e int) {
	output := rpc.FunctionOutput{
		Error:  e,
		Result: nil,
		Meta:   nil,
	}

	data, err := json.Marshal(output)
	if err != nil {
		return
	}

	_, _ = writer.Write(data)
}
