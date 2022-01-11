package rpc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var client = NewClient("http://testrpc.cadmean.ru", DefaultConfiguration())

func TestFunction_Call(t *testing.T) {
	sum, err := client.Func("sum").Call(2, 67)

	if err != nil {
		t.Error(err)
	}

	expected := 69.0

	assert.Equal(t, expected, sum.Result)
}

func TestFunction_CallForResult(t *testing.T) {
	var sum float64
	err := client.Func("sum").CallForResult(&sum, 2, 67)
	if err != nil {
		t.Error(err)
	}

	var expected float64 = 69

	assert.Equal(t, expected, sum)
}

func Test_CallConcat(t *testing.T) {
	var concatStr string
	err := client.Func("concat").CallForResult(&concatStr, "Hello,", " RPC!")
	if err != nil {
		t.Error(err)
	}

	var expected = "Hello, RPC!"

	assert.Equal(t, expected, concatStr)
}

type User struct {
	Name, Email string
}

func Test_CallUserGet(t *testing.T) {
	_, err := client.Func("auth").Call("email@example.com", "password")
	if err != nil {
		t.Error(err)
	}
	assert.NotEmpty(t, client.configuration.AuthTicketHolder.GetTicket().AccessToken)

	var user User
	err = client.Func("user.get").CallForResult(&user)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "George", user.Name)
	t.Logf("%+v", user)
}
