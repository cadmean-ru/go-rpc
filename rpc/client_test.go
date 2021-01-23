package rpc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var client = NewClient("http://testrpc.cadmean.ru")

func TestFunction_Call(t *testing.T) {
	sum, err := client.F("sum").Call(2, 67)

	if err != nil {
		t.Error(err)
	}

	expected := 69.0

	assert.Equal(t, expected, sum)
}
