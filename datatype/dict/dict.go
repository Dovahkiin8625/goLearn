package dict

import "fmt"

func DictStatement() {
	// 声明并初始化
	var dict1 map[string]int = map[string]int{"age": 16, "year": 2022}
	dict2 := map[string]int{"age": 16, "year": 2022}
	dict3 := make(map[string]int)
	dict3["age"] = 16
	dict3["year"] = 2022
	fmt.Println(dict1)
	fmt.Println(dict2)
	fmt.Println(dict3)
}

// 字典相关操作
func DictOper() {
	dict1 := map[string]int{"age": 16, "year": 2022}
	// 添加,更新元素
	dict1["month"] = 10
	dict1["year"] = 2021
	fmt.Println()
	fmt.Println(dict1)
	// 读取元素
	fmt.Println(dict1["month"])
	fmt.Println(dict1["day"])
	// 删除元素
	delete(dict1, "month")
	fmt.Println(dict1)
}

// 判断key 存在
func DictKeyExists() {
	dict1 := map[string]int{"age": 17, "year": 2022}
	checkKeyExists := func(dict map[string]int, key string) {
		if value, exist := dict[key]; exist {
			fmt.Printf("key: %s 存在,值为: %d\n", key, value)
		} else {
			fmt.Printf("key: %s 不存在\n", key)
		}
	}
	fmt.Println()
	checkKeyExists(dict1, "month")
	checkKeyExists(dict1, "year")
}

// 字典遍历
func DictRange() {
	dict1 := map[string]int{"age": 16, "year": 2022, "month": 2, "day": 1, "hour": 12}
	fmt.Println()
	for key, value := range dict1 {
		fmt.Printf("key:%s,value:%d\n", key, value)
	}
	for key := range dict1 {
		fmt.Printf("key:%s\n", key)
	}
	for _, value := range dict1 {
		fmt.Printf("value:%d\n", value)
	}
}
