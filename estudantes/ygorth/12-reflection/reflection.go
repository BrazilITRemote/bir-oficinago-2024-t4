package reflection

import (
	"fmt"
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	tipo := val.Type()
	fmt.Printf("\n Tipo de valor: %s \n", tipo)

	// if val.Kind() != reflect.Struct {
	// 	return
	// }

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String {
			fn(field.String())
		}
	}
}
