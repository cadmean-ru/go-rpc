package rpc

import (
	"errors"
	"fmt"
	"github.com/cadmean-ru/require"
	"github.com/mitchellh/mapstructure"
	"reflect"
)

type Function struct {
	name   string
	client *Client
}

func (f *Function) Call(args ...interface{}) (*FunctionOutput, error) {
	conf := f.client.configuration

	call := FunctionCall{
		Arguments: args,
		Auth:      conf.AuthTicketHolder.GetTicket().AccessToken,
	}

	url := fmt.Sprintf("%s/%s", f.client.url, conf.FunctionUrlProvider.GetUrl(f.name))

	callData, err := conf.CodecProvider.Encode(call)
	if err != nil {
		return nil, NewError(ErrorEncode, "failed to encode call")
	}

	outputData, err := conf.TransportProvider.Send(url, callData, conf.CodecProvider.ContentType())
	if err != nil {
		return nil, err
	}

	output := &FunctionOutput{}
	err = conf.CodecProvider.Decode(outputData, output)
	if err != nil {
		return nil, NewError(ErrorDecode, "failed to decode response")
	}

	if output.Error != 0 {
		return output, NewError(output.Error, fmt.Sprintf("function call finished with error %d", output.Error))
	}

	if resultType, ok := output.Meta["resultType"]; ok && resultType == "auth" {
		m := require.SiMap(output.Result)
		ticket := NewAuthTicket(require.String(m["accessToken"]), require.String(m["refreshToken"]))
		f.client.configuration.AuthTicketHolder.SetTicket(ticket)
	}

	return output, nil
}

func (f *Function) CallForResult(result interface{}, args ...interface{}) error {
	output, err := f.Call(args...)
	if err != nil {
		return err
	}

	resultType := reflect.TypeOf(result)

	if resultType.Kind() != reflect.Ptr {
		return errors.New("not a pointer")
	}

	resultType = resultType.Elem()
	resultValue := reflect.ValueOf(result)
	resultValue = reflect.Indirect(resultValue)
	if resultType.Kind() == reflect.Map || resultType.Kind() == reflect.Struct {
		err = mapstructure.Decode(output.Result, result)
		if err != nil {
			return err
		}
	} else {
		resultValue.Set(reflect.ValueOf(output.Result))
	}

	return nil
}

func newFunction(name string, client *Client) *Function {
	return &Function{
		name:   name,
		client: client,
	}
}
