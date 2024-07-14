package bank

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_EmptyPrint(t *testing.T) {
	account := new(StandardAccount)
	actual := account.printStatement()
	expected := "Date       || Amount || Balance"
	assert.Equal(t, expected, actual, "The headers are not the same")
}

func Test_SingleDeposit(t *testing.T) {
	account := new(StandardAccount)

	account.deposit(100)
	actual := account.printStatement()
	expected := "" +
		"Date       || Amount || Balance\n" +
		"2024-07-14 || 100    || 100"

	assert.Equal(t, expected, actual, "Deposit should be there")
}
