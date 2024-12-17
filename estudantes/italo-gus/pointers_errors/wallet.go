package pointers_erros // Ponteiros e Erros
import "fmt"

type Wallet struct {
	balance int
}

/*
func (w Wallet) Deposit(amount int) {
	fmt.Printf("memory address of balance in Deposit is %p \n", &w.balance)
	// & antes da variável é seu endereço na memória em hexadecimal
	// %p notação base 16, com 0x à esquerda
	w.balance += amount
}

func (w Wallet) Balance() int {
	fmt.Printf("memory address of balance in Balance is %p \n", &w.balance)
	return w.balance
}
*/

// Em Go copia valores quando você os passa para funções/métodos como argumento,
// então se você estiver escrevendo uma função que precisa alterar o estado,
// você precisará que ela receba um ponteiro para o que você quer alterar.
// https://go.dev/ref/spec#Method_values
// https://go.dev/doc/effective_go#pointers_vs_values
// https://go.dev/ref/spec#Pointer_types
// https://gobyexample.com/pointers
// https://gobyexample.com/methods

func (w *Wallet) Deposit(amount int) {
	fmt.Printf("memory address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() int {
	fmt.Printf("memory address of balance in Balance is %p \n", &w.balance)
	return w.balance
}
