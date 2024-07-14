package bank

import (
	"testing"
)

func Test_EmptyPrint(t *testing.T) {
	account := StandardAccount{}
	actual := account.printStatement()
	expected := "Date       || Amount || Balance"
	if actual != expected {
		t.Errorf("The headers are not as expected.\n%s\nvs\n%s", actual, expected)
	}
}
