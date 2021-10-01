package rpc

type Configuration struct {
	TransportProvider   TransportProvider
	CodecProvider       CodecProvider
	FunctionUrlProvider FunctionUrlProvider
	AuthTicketHolder    AuthTicketHolder
}

func DefaultConfiguration() *Configuration {
	return &Configuration{
		TransportProvider:   NewHttpTransportProvider(),
		CodecProvider:       NewJsonCodecProvider(),
		FunctionUrlProvider: NewDefaultFunctionUrlProvider(),
		AuthTicketHolder:    NewTransientAuthTicketHolder(),
	}
}