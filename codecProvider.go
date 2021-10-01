package rpc

import "encoding/json"

type CodecProvider interface {
	ContentType() string
	Encode(source interface{}) ([]byte, error)
	Decode(data []byte, destination interface{}) error
}

type JsonCodecProvider struct {

}

func (j *JsonCodecProvider) ContentType() string {
	return "application/json"
}

func (j *JsonCodecProvider) Encode(source interface{}) ([]byte, error) {
	return json.Marshal(source)
}

func (j *JsonCodecProvider) Decode(data []byte, destination interface{}) error {
	return json.Unmarshal(data, destination)
}

func NewJsonCodecProvider() *JsonCodecProvider {
	return &JsonCodecProvider{}
}