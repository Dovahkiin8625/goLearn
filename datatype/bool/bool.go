package bool

import "fmt"

func TestBool() {
	btrue := true
	bfalse := false
	fmt.Println()
	// bool 转 int
	bool2Int := func(b bool) int {
		if b {
			return 1
		}
		return 0
	}
	fmt.Println(bool2Int(btrue))
	fmt.Println(bool2Int(bfalse))
	// int 转bool
	int2Bool := func(i int) bool {
		return i != 0
	}
	fmt.Println(int2Bool(1))
	fmt.Println(int2Bool(0))
}
