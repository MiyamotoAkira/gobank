package bank

import (
	"strconv"
	"time"
)

type Account interface {
	deposit(amount int)
	withdraw(amount int)
	printStatement() string
}

type Statement struct {
	Date    string
	Amount  int
	Balance int
}

type StandardAccount struct {
	statements []Statement
}

func (account *StandardAccount) printStatement() string {
	result := "Date       || Amount || Balance"

	for _, statement := range account.statements {
		result += "\n" + statement.Date + " || " + strconv.Itoa(statement.Amount) + "    || " + strconv.Itoa(statement.Balance)
	}

	return result
}

func (account *StandardAccount) deposit(amount int) {
	date := time.Now().Format("2006-01-02")

	current_balance := 0
	if len(account.statements) > 0 {
		current_balance = account.statements[0].Balance
	}
	statement := Statement{date, amount, amount + current_balance}
	account.statements = append([]Statement{statement}, account.statements...)
}

func (account *StandardAccount) withdraw(amount int) {
	date := time.Now().Format("2006-01-02")

	current_balance := 0
	if len(account.statements) > 0 {
		current_balance = account.statements[0].Balance
	}
	statement := Statement{date, -amount, current_balance - amount}
	account.statements = append([]Statement{statement}, account.statements...)
}
