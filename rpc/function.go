package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Function struct {
	name string
	client *Client
}

func (f *Function) Call(args ...interface{}) (interface{}, error) {
	call := FunctionCall{
		Arguments: args,
	}

	url := fmt.Sprintf("%s/api/rpc/%s", f.client.url, f.name)

	jsonStr, err := json.Marshal(call)
	if err != nil {
		return nil, NewError(ErrorEncode, "failed to encode call")
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, NewError(ErrorEncode, "failed to encode call")
	}

	req.Header.Set("Cadmean-RPC-Version", "2.1")
	req.Header.Set("Content-Type", "application/json")

	resp, err := f.client.httpClient.Do(req)
	if err != nil {
		return nil, NewError(ErrorCouldNotSendCall, "http error")
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, NewError(ErrorNotSuccessfulStatusCode, "not successful status code")
	}

	body, _ := ioutil.ReadAll(resp.Body)

	output := FunctionOutput{}
	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, NewError(ErrorDecode, "failed to decode response")
	}

	if output.Error != 0 {
		return nil, NewError(output.Error, fmt.Sprintf("Function call finished with error %d", output.Error))
	}

	return output.Result, nil
}

func newFunction(name string, client *Client) *Function {
	return &Function{
		name: name,
		client: client,
	}
}