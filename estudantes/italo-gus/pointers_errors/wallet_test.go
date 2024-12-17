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
/*
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

*/

/*
func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		received := wallet.Balance()
		fmt.Printf("memory address of balance in test deposit is %p \n", &wallet.balance)
		expected := Bitcoin(10)
		if received != expected {
			t.Errorf("expected '%s' but received '%s'", expected, received)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		received := wallet.Balance()
		fmt.Printf("memory address of balance in test withdraw is %p \n", &wallet.balance)
		expected := Bitcoin(10)
		if received != expected {
			t.Errorf("expected '%s' but received '%s'", expected, received)
		}
	})
}
*/

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, expected Bitcoin) {
		t.Helper()
		received := wallet.Balance()

		if received != expected {
			t.Errorf("expected '%s' but received '%s'", expected, received)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		fmt.Printf("memory address of balance in test deposit is %p \n", &wallet.balance)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)} // criação já estipulando valor de balance
		fmt.Printf("memory address of balance in test withdraw is %p \n", &wallet.balance)
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

}
