package main

import (
	"fmt"

	"vancode.top/goLearn/datatype/byterunestring"
	"vancode.top/goLearn/datatype/float"
	"vancode.top/goLearn/datatype/int"
	"vancode.top/goLearn/datatype/list"
)

func divideLine(line string) {
	fmt.Println()
	fmt.Println(line)
	fmt.Println("--------------------------------")
}

// int类型测试
func testInt() {
	divideLine("Test go int data type >>")
	int.SizeIntUnit()
	int.PrintBinary()
	int.PrintOctal()
	int.PrintHexadecimal()
}

// 浮点类型测试
func testFloat() {
	divideLine("Test go float data type >>")
	float.PrintMax_Min()
}

// 字符字符串类型测试
func testString() {
	divideLine("Test go string data type >>")
	byterunestring.PrintByteUint()
	byterunestring.PrintString()
}

// 列表测试
func testList() {
	divideLine("Test go list data type >>")
	list.ListStatement()
	list.ListAutoSize()
}
func main() {
	testInt()
	testFloat()
	testString()
	testList()
}
