package main

import (
	"fmt"
)

var (
	name string = "linhao"
	age  int    = 123
)

const (
	NAME string = "LINHAO"
	AGE  int    = 12
)

func printVar() {
	name := "Vance"
	fmt.Println(name)
}
func mutiVar() {
	name, age := "Vance", 24
	fmt.Println(name)
	fmt.Println(age)
}
func ptrVar() {
	ptr := new(int)
	fmt.Println(ptr)
	fmt.Println(*ptr)
}
func ptrInt() *int {
	return new(int)
}
func ptrIntNew() *int {
	var tempInt int
	return &tempInt
}
func main() {
	// fmt.Println(name)
	// fmt.Println(age)
	// fmt.Println(NAME)
	// fmt.Println(AGE)
	printVar()
	mutiVar()
	ptrVar()
}
