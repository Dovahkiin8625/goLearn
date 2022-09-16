package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Head string
	Type string
}

type AnimalTag struct {
	Head string `label:"Head"`
	Type string `label:"Type"`
}

func printField(typeR reflect.Type) {
	field := typeR.Field(0)
	fmt.Print("Get field by index: ")
	fmt.Println(field)

	field1, hasF := typeR.FieldByName("Head")
	if !hasF {
		fmt.Println("Field 'Head' not found")
	} else {
		fmt.Print("Get field by name: ")
		fmt.Println(field1)
	}
	fmt.Println()
}

func main() {
	// 三种方式获取field
	// 方法1:
	typeR := reflect.TypeOf(AnimalTag{})
	typeRv := reflect.ValueOf(AnimalTag{}).Type()
	typeRp := reflect.ValueOf(&AnimalTag{}).Elem().Type()

	printField(typeR)
	printField(typeRv)
	printField(typeRp)

}
