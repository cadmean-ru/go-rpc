package rpc

import "github.com/cadmean-ru/require"

type AuthTicket struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (a *AuthTicket) FromMap(m map[string]interface{}) {
	a.AccessToken = require.String(m["accessToken"])
	a.RefreshToken = require.String(m["refreshToken"])
}

func NewAuthTicket(access, refresh string) AuthTicket {
	return AuthTicket{
		AccessToken:  access,
		RefreshToken: refresh,
	}
}
