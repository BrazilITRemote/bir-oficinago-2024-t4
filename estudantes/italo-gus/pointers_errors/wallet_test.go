package pointers_erros // Delcaração de tipo de dados

import (
	"fmt"
	"testing"
)

/*
func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)
	received := wallet.Balance()

	fmt.Printf("memory address of balance in test is %p \n", &wallet.balance)

	expected := 10
	if received != expected {
		t.Errorf("expected '%d' but received '%d'", expected, received)
	}
}
*/

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))
	received := wallet.Balance()

	fmt.Printf("memory address of balance in test is %p \n", &wallet.balance)

	expected := Bitcoin(10)
	if received != expected {
		t.Errorf("expected '%s' but received '%s'", expected, received)
	}
}
