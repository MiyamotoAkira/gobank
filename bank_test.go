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

func Test_DepositsOnDifferentDates(t *testing.T) {
	mockedDate := Date("")
	mockedDateProvider := func() Date {
		return mockedDate
	}
	account := StandardAccount{TodaysDateProvider: mockedDateProvider}

	mockedDate = "2012-01-10"
	account.deposit(1000)
	mockedDate = "2012-01-13"
	account.deposit(2000)
	mockedDate = "2012-01-14"
	account.withdraw(500)

	expected := "" +
		"Date       || Amount || Balance\n" +
		"2012-01-14 || -500    || 2500\n" +
		"2012-01-13 || 2000    || 3000\n" +
		"2012-01-10 || 1000    || 1000"
	actual := account.printStatement()
	assert.Equal(t, expected, actual, "Expected statement should match")
}
