package iteration

import "strings"

// Returns a new string, by repeating a character n times
func Repeat(character string, n_times int) string {
	// var repeated string
	// for i := 0; i < n_times; i++ {
	// 	repeated += character
	// }
	// return repeated
	return strings.Repeat(character, n_times)
}
