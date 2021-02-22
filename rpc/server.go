package rpc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const Version = "2.1"

type HandlerFunc func(args ...interface{}) (interface{}, error)

func MakeHandler(delegate HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		v := request.Header.Get("Cadmean-RPC-Version")
		if v != Version {
			writeError(writer, ErrorIncompatibleRPCVersion)
			return
		}

		writer.Header().Add("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS")
		writer.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		defer request.Body.Close()
		data, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writeError(writer, ErrorDecode)
			return
		}

		call := FunctionCall{}
		err = json.Unmarshal(data, &call)
		if err != nil {
			writeError(writer, ErrorDecode)
			return
		}

		defer func() {
			recoveredError := recover()
			if recoveredError != nil {
				writeError(writer, ErrorServer)
			}
		}()

		result, customError := delegate(call.Arguments...)
		var errorCode int
		if customError != nil {
			if rpcError, ok := customError.(Error); ok {
				errorCode = rpcError.Code
			} else {
				errorCode = ErrorServer
			}
		}

		output := FunctionOutput{
			Error:  errorCode,
			Result: result,
			Meta:   nil,
		}

		responseData, err := json.Marshal(output)
		if err != nil {
			writeError(writer, ErrorEncode)
			return
		}

		_, _ = writer.Write(responseData)
	}
}

func Handle(functionName string, handler HandlerFunc) {
	http.HandleFunc(fmt.Sprintf("/api/rpc/%s", functionName), MakeHandler(handler))
}

func writeError(writer http.ResponseWriter, e int) {
	output := FunctionOutput{
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