package test

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	total, desc := useShortParam(3, 5, "求和")
	//total, desc := noUseShortParam(3, 5, "求和")

	fmt.Printf("total: %d desc: %s", total, desc)

	uncertain(1,2,3,4,5,6)
}

//参数不简写，并且使用返回：变量名+变量类型组合
func noUseShortParam(a int, b int, str string) (total int, desc string) {
	//返回中已经声明返回变量名称，这里可以直接使用
	total = a + b
	desc = fmt.Sprintf("%s: %d + %d = %d \n", str, a, b, total)
	return total, desc
}

func useShortParam(a, b int, str string) (int, string) {
	//返回中只声明返回变量类型，这里需要声明变量才能使用
	total := a + b
	desc := fmt.Sprintf("%s: %d + %d = %d \n", str, a, b, total)
	return total, desc
}

// arg ...type：告诉函数接收不定数量的参数。在函数体中，变量arg是一个int的slice
/*
注意事项：
	一个函数最多只能有一个可变参数
	若参数列表中还有其他类型参数，则可变参数写在所有参数的最后
*/
func uncertain(number ...int) int {
	fmt.Printf("类型：%T 值: %v \n", number, number)
	var total int
	for _, val := range number {
		total += val
	}
	fmt.Printf("总和: %v", total)
	return total
}
