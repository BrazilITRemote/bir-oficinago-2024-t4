package iteration // Iteração / Laços de Repetições

func Repeat(character string) string {
	/*	Em Go não há parênteses em torno dos três componentes da instrução for
		Em Go não possui nenhuma palavra chave do tipo while, do ou until.
		https://go.dev/doc/effective_go#for
		https://go.dev/ref/spec#For_statements
	*/
	var repeated string = ""
	for i := 0; i < 5; i++ {
		/* operandor incrementam ++ ou decrementam --
		https://go.dev/ref/spec#IncDec_statements */
		repeated = repeated + character
	}
	return repeated
}
