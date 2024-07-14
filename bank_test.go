package bank

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func MockedDateProvider() Date {
	return "2024-07-14"
}

func Test_EmptyPrint(t *testing.T) {
	account := StandardAccount{TodaysDateProvider: MockedDateProvider}
	actual := account.printStatement()
	expected := "Date       || Amount || Balance"
	assert.Equal(t, expected, actual, "The headers are not the same")
}

func Test_SingleDeposit(t *testing.T) {
	account := StandardAccount{TodaysDateProvider: MockedDateProvider}

	account.deposit(100)
	actual := account.printStatement()
	expected := "" +
		"Date       || Amount || Balance\n" +
		"2024-07-14 || 100    || 100"

	assert.Equal(t, expected, actual, "Deposit should be there")
}

func Test_TwoDeposit(t *testing.T) {
	account := StandardAccount{TodaysDateProvider: MockedDateProvider}

	account.deposit(100)
	account.deposit(200)
	actual := account.printStatement()
	expected := "" +
		"Date       || Amount || Balance\n" +
		"2024-07-14 || 200    || 300\n" +
		"2024-07-14 || 100    || 100"

	assert.Equal(t, expected, actual, "Deposit should be there")
}

func Test_DepositThenWithdrawal(t *testing.T) {
	account := StandardAccount{TodaysDateProvider: MockedDateProvider}

	account.deposit(200)
	account.withdraw(100)
	actual := account.printStatement()
	expected := "" +
		"Date       || Amount || Balance\n" +
		"2024-07-14 || -100    || 100\n" +
		"2024-07-14 || 200    || 200"

	assert.Equal(t, expected, actual, "Deposit should be there")
}
