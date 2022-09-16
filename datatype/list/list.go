package list

import "fmt"

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
	arr2 := []int{1, 2, 3, 4}
	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(len(arr))
	fmt.Println(len(arr2))
}

// 数组类型
func ListType() {

}
