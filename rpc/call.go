package rpc

type FunctionCall struct {
	Arguments []interface{} `json:"args"`
	Auth      string        `json:"auth"`
}
