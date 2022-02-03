package rpc

import "fmt"

type FunctionUrlProvider interface {
	GetUrl(functionName string) string
}

type DefaultFunctionUrlProvider struct {
	Prefix string
}

func (p *DefaultFunctionUrlProvider) GetUrl(functionName string) string {
	return fmt.Sprintf("%s/%s", p.Prefix, functionName)
}

func NewDefaultFunctionUrlProvider() *DefaultFunctionUrlProvider {
	return &DefaultFunctionUrlProvider{Prefix: "api/rpc"}
}