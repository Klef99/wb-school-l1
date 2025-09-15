package main

import (
	"fmt"
	"reflect"
)

func getType(v interface{}) string {
	t := reflect.TypeOf(v)
	if t == nil {
		return "unknown"
	}
	// для определения канала не прибегая к конкретному типу (chan int и т.д.),
	// пришлось использовать reflect вместо switch.(type)
	switch t.Kind() {
	case reflect.Int:
		return "int"
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "bool"
	case reflect.Chan:
		return "chan"
	default:
		return "unknown"
	}
}

func main() {
	ch := make(chan int)
	i := 1
	s := "str"
	b := true
	fmt.Println(getType(i), getType(b), getType(ch), getType(s))
}
