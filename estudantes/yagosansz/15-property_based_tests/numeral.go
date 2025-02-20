package main

import "strings"

func ConvertToRoman(arabic int) string {
	// A Builder is used to efficiently build a string using Write methods.
	// It minimizes memory copying.
	var result strings.Builder

	for i := 0; i < arabic; i++ {
		result.WriteString("I")
	}

	return result.String()
}
