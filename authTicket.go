package rpc

type AuthTicket struct {
	AccessToken, RefreshToken string
}

func NewAuthTicket(access, refresh string) AuthTicket {
	return AuthTicket{
		AccessToken: access,
		RefreshToken: refresh,
	}
}