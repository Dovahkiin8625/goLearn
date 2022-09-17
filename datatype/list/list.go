package list

import (
	"fmt"
	"reflect"
)

// 数组声明
func ListStatement() {
	// 声明
	var arr1 [3]int
	// 声明并初始化
	var arr2 [3]int = [3]int{1, 2, 3}
	// 初始化声明
	arr3 := [3]int{1, 2, 3}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}

// 数组自动分配空间
func ListAutoSize() {
	arr := [...]int{1, 2, 3}
	fmt.Println()
	fmt.Println(arr)
	fmt.Println(len(arr))
}

// 数组类型
func ListType() {

	arr1 := [...]int{1, 2, 3}
	arr2 := [...]int{1, 2, 3, 4}
	fmt.Println()
	// %T 输出变量类型
	fmt.Printf("%T\n", arr1)
	fmt.Printf("%T\n", arr2)
	// 反射输出变量类型
	fmt.Println(reflect.TypeOf(arr1))
	fmt.Println(reflect.TypeOf(arr2))
}

// 数组类型别名
func ListTypeAlias() {
	type arrType [3]int
	myarr := arrType{1, 2, 3}
	fmt.Println()
	fmt.Println(myarr)
	fmt.Println(reflect.TypeOf(myarr))
}

// 按索引初始化
func ListIndex() {
	arr := [4]int{2: 3}
	fmt.Println()
	fmt.Println(arr)
}
