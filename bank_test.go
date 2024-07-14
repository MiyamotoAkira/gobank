package bank

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_EmptyPrint(t *testing.T) {
	account := StandardAccount{}
	actual := account.printStatement()
	expected := "Date       || Amount || Balance"
	assert.Equal(t, expected, actual, "The headers are not the same")
}
