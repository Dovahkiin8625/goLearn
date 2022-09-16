package byterunestring

import "fmt"

func PrintByteUint() {
	var a byte = 65
	var b uint = 66
	var a2 byte = 'A'
	var b2 uint = 'B'
	fmt.Printf("a = %c, b = %c\n", a, b)
	fmt.Printf("a2 = %c, b2 = %c\n", a2, b2)
}
func PrintString() {
	var str1 string = "hello \\world"
	var str2 string = `hello \world`
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Printf("%q\n", str1)
	fmt.Print(`hello
world`)
	fmt.Println()
}
