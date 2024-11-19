package iteration // Refactor

const repeatCount = 5

func Repeat(character string) string {
	var repeated string = ""
	for i := 0; i < repeatCount; i++ {
		repeated += character
		/* += Concatenação de strings
		https://go.dev/ref/spec#String_concatenation */
	}
	return repeated
}
