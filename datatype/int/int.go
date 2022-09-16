package int

import (
	"fmt"
	"unsafe"
)

func SizeIntUnit() {
	var intVar int
	var uintVar uint
	fmt.Printf("intVar 数据类型为 %T,大小为 %d\n", intVar, unsafe.Sizeof(intVar))
	fmt.Printf("uintVar 数据类型为 %T,大小为 %d\n", uintVar, unsafe.Sizeof(uintVar))
}
func PrintBinary() {
	num := 0b1100
	fmt.Printf("二进制数%b,十进制数%d\n", num, num)
}
func PrintOctal() {
	num := 0o1100
	fmt.Printf("八进制数%o,十进制数%d\n", num, num)
}
func PrintHexadecimal() {
	num := 0x1100
	fmt.Printf("十六进制数%X,十进制数%d\n", num, num)
}
