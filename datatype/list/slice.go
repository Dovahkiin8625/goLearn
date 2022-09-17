package list

import "fmt"

// 从数组截取
func SliceCutfromList() {
	list := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	// 表示截取数组list的从索引1开始,到索引2(3-1)截至的元素
	slice1 := list[1:3]
	// 设置切片容量的截止索引
	slice2 := list[1:3:4]
	fmt.Printf("切片 slice1: %d 的长度为: %d,容量为: %d \n", slice1, len(slice1), cap(slice1))
	fmt.Printf("切片 slice2: %d 的长度为: %d,容量为: %d \n", slice2, len(slice2), cap(slice2))
}

// 切片声明
func SliceStatement() {
	var slice1 []int
	slice1 = append(slice1, 123)
	var slice2 []int = []int{1, 2, 3, 4}
	slice3 := []int{1, 2, 3, 4}
	fmt.Println()
	fmt.Printf("切片 slice1: %d 的长度为: %d,容量为: %d \n", slice1, len(slice1), cap(slice1))
	fmt.Printf("切片 slice2: %d 的长度为: %d,容量为: %d \n", slice2, len(slice2), cap(slice2))
	fmt.Printf("切片 slice3: %d 的长度为: %d,容量为: %d \n", slice3, len(slice3), cap(slice3))
}

// make函数构造
func SliceMake() {
	slice1 := make([]int, 2)
	slice2 := make([]int, 2, 6)
	fmt.Println()
	fmt.Printf("切片 slice1: %d 的长度为: %d,容量为: %d \n", slice1, len(slice1), cap(slice1))
	fmt.Printf("切片 slice2: %d 的长度为: %d,容量为: %d \n", slice2, len(slice2), cap(slice2))
	// fmt.Printf("切片 slice3: %d 的长度为: %d,容量为: %d \n", slice3, len(slice3), cap(slice3))
}

// 指定索引对应值构造
func SliceIndex() {
	slice1 := []int{2: 13}
	fmt.Println()
	fmt.Printf("切片 slice1: %d 的长度为: %d,容量为: %d \n", slice1, len(slice1), cap(slice1))
}

//
