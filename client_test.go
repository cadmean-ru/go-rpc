package rpc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var client = NewClient("http://testrpc.cadmean.ru", DefaultConfiguration())

func TestFunction_Call(t *testing.T) {
	sum, err := client.F("sum").Call(2, 67)

	if err != nil {
		t.Error(err)
	}

	expected := 69.0

	assert.Equal(t, expected, sum.Result)
}

func TestFunction_CallForResult(t *testing.T) {
	var sum float64
	err := client.F("sum").CallForResult(&sum, 2, 67)
	if err != nil {
		t.Error(err)
	}

	var expected float64 = 69

	assert.Equal(t, expected, sum)
}