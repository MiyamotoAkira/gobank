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
		result += "\n" + statement.Date + " || " + strconv.Itoa(statement.Amount) + "   || " + strconv.Itoa(statement.Balance)
	}

	return result
}

func (account *StandardAccount) deposit(amount int) {
	date := time.Now().Format("2006-01-02")

	account.statements = append(account.statements, Statement{date, amount, amount})
}
