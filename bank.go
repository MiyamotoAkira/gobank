package bank

import (
	"strconv"
)

type Account interface {
	deposit(amount int)
	withdraw(amount int)
	printStatement() string
}

type Statement struct {
	Date    Date
	Amount  int
	Balance int
}

type StandardAccount struct {
	statements         []Statement
	TodaysDateProvider TodaysDateProvider
}

func NewStandardAccount() *StandardAccount {
	return &StandardAccount{TodaysDateProvider: TodaysSystemDateProvider}
}

func (account *StandardAccount) printStatement() string {
	result := "Date       || Amount || Balance"

	for _, statement := range account.statements {
		result += "\n" + string(statement.Date) + " || " + strconv.Itoa(statement.Amount) + "    || " + strconv.Itoa(statement.Balance)
	}

	return result
}

func (account *StandardAccount) deposit(amount int) {
	date := account.TodaysDateProvider()

	currentBalance := 0
	if len(account.statements) > 0 {
		currentBalance = account.statements[0].Balance
	}
	statement := Statement{date, amount, amount + currentBalance}
	account.statements = append([]Statement{statement}, account.statements...)
}

func (account *StandardAccount) withdraw(amount int) {
	date := account.TodaysDateProvider()

	currentBalance := 0
	if len(account.statements) > 0 {
		currentBalance = account.statements[0].Balance
	}
	statement := Statement{date, -amount, currentBalance - amount}
	account.statements = append([]Statement{statement}, account.statements...)
}
