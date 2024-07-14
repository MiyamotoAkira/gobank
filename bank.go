package bank

type Account interface {
	deposit(amount int)
	withdraw(amount int)
	printStatement() string
}

type StandardAccount struct {
}

func (a StandardAccount) printStatement() string { return "Date       || Amount || Balance" }
