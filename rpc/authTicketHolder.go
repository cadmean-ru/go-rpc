package rpc

type AuthTicketHolder interface {
	SetTicket(ticket AuthTicket)
	GetTicket() AuthTicket
}

type TransientAuthTicketHolder struct {
	ticket AuthTicket
}

func (t *TransientAuthTicketHolder) SetTicket(ticket AuthTicket) {
	t.ticket = ticket
}

func (t *TransientAuthTicketHolder) GetTicket() AuthTicket {
	return t.ticket
}

func NewTransientAuthTicketHolder() *TransientAuthTicketHolder {
	return &TransientAuthTicketHolder{}
}