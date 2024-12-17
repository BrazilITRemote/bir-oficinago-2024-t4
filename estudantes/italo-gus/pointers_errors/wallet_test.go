package pointers_erros // Delcaração de tipo de dados

import (
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

/*
func TestWallet(t *testing.T) {
	// criando função como variável para tratamento de erro
	assertError := func(t testing.TB, received error, expected string) {
		t.Helper()

		if received == nil {
			t.Fatal("didn't get an error but wanted one")
			// t.Fatal exibe o texto em seguida chama t.FatalNow
			// t.FatalNow marca a função como tendo falhado e
			// interrompe sua execução chamando runtime.Goexit
			// (que então executa todas as chamadas adiadas na goroutine atual).
			// A execução continuará no próximo teste ou benchmark.
			// https://pkg.go.dev/testing#T.Fatal
			// https://pkg.go.dev/github.com/cszatmary/goutils/fatal

		}

		// convertendo erro em string com a função Error()
		// Error() é uma interface interna semelhante a fmt.Stringer
		// https://go.dev/blog/error-handling-and-go
		// https://go.dev/doc/tutorial/handle-errors
		// https://goporexemplo.golangbr.org/errors.html
		// https://go.dev/tour/methods/19

		if received.Error() != expected {
			t.Errorf("expected '%q' but received '%q'", expected, received)
		}
	}

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
		wallet := Wallet{balance: Bitcoin(20)} // criação já estipulando valor de balance 1
		fmt.Printf("memory address of balance in test withdraw is %p \n", &wallet.balance)
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance} // criação já estipulando valor de balance 2
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, "cannot withdraw, insufficient funds")
		assertBalance(t, wallet, startingBalance)

		// tratando o  erro  de forma simples
		//	if err == nil {
		//		t.Error("wanted an error but didn't get one")
		//	}


	})
}
*/

/*
func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		// sem tratar o erro de retorno
		// para varredura utiliza um linters
		// Um dos muitos linters disponíveis para Go é o errcheck
		// https://github.com/kisielk/errcheck
		// instalação
		// go install github.com/kisielk/errcheck@latest
		// para usar dentro do diretório com seu código, execute errcheck
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(20))
	})
}

func assertBalance(t testing.TB, wallet Wallet, expected Bitcoin) {
	t.Helper()
	received := wallet.Balance()
	if received != expected {
		t.Errorf("expected '%q' but received '%q'", expected, received)
	}
}

// criando função para tratamento de erro
func assertError(t testing.TB, received, expected error) {
	t.Helper()
	if received == nil {
		t.Fatal("didn't get an error but wanted one")
	}
	if received != expected {
		t.Errorf("expected '%q' but received '%q'", expected, received)
	}
}
*/

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		// assertNoError para tratar o erro, mesmo prevendo nil,
		// como não há como descarta com _ blank identifier.
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(20))
	})
}

func assertBalance(t testing.TB, wallet Wallet, expected Bitcoin) {
	t.Helper()
	received := wallet.Balance()
	if received != expected {
		t.Errorf("expected '%s' but received '%s'", expected, received)
	}
}

// função para tratar os erros prevendo nil, devido não ter como descarta com _ blank identifier.
func assertNoError(t testing.TB, received error) {
	t.Helper()
	if received != nil {
		t.Fatal("got an error but didn't want one")
	}
}

// Função para tratar os erros
func assertError(t testing.TB, received error, expected error) {
	t.Helper()
	if received == nil {
		t.Fatal("didn't get an error but wanted one")
	}
	if received != expected {
		t.Errorf("expected '%s' but received '%s'", expected, received)
	}
}
