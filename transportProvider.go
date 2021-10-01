package rpc

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TransportProvider interface {
	Send(url string, data []byte, contentType string) ([]byte, error)
}

type HttpTransportProvider struct {
	*http.Client
}

func (p *HttpTransportProvider) Send(url string, data []byte, contentType string) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, NewError(ErrorCouldNotSendCall, "failed to create request")
	}

	req.Header.Set("Cadmean-RPC-Version", Version)
	req.Header.Set("Content-Type", contentType)

	resp, err := p.Do(req)
	if err != nil {
		return nil, NewError(ErrorCouldNotSendCall, "failed to send request")
	}
	if resp.StatusCode != 200 {
		return nil, NewError(ErrorNotSuccessfulStatusCode, fmt.Sprintf("response status code: %d", resp.StatusCode))
	}

	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, NewError(ErrorCouldNotSendCall, "failed to read response")
	}

	return respBytes, nil
}

func NewHttpTransportProvider() *HttpTransportProvider {
	return &HttpTransportProvider{Client: &http.Client{}}
}