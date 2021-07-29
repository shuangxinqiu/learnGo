package test

import (
	"fmt"
	"testing"
)

func TestForRange(t *testing.T) {
	//	遍历字符串
	str := "hello,welcome to"
	for key, val := range str {
		fmt.Printf("第%d位,字符是: %c ASCII值为：%d \n", key, val, val)
	}

	// 遍历数组
	arr := [6]int{1, 2, 3, 4, 5}
	fmt.Printf("arr 类型: %T, 值:%v \n", arr, arr)
	for key2, val2 := range arr {
		fmt.Printf("key: %d, value: %d \n", key2, val2)
	}

	//遍历切片
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	// key 可以省略不写
	for val3 := range slice {
		fmt.Printf("值: %d \n", val3)
	}

	//遍历Map

	map2 := make(map[string]int)
	map2["a"] = 1
	map2["v"] = 2
	for key4, val4 := range map2 {
		fmt.Printf("key: %s , val: %d \n", key4, val4)
	}
}
