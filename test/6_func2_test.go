package test

import (
	"fmt"
	"strings"
	"testing"
)

type strConvert func(string) string

func TestFunc2(t *testing.T) {
	//函数类型，函数也是类型，可以和int32，float等一样被保存在变量中
	//在Go语言中可以通过type定义一个自定义类型。函数的参数完全相同，返回值相同
	res := testFunc("hello,world, go is best language", funcParam)
	fmt.Println(res)

	res2 := testFunc2("hello,world, go is best language", funcParam)
	fmt.Println(res2)
}

//不声明函数类型
//定义一个函数，接收字符串和函数类型：f func(string) string
func testFunc(str string, f func(string) string) string {
	return f(str)
}

//声明函数类型
//定义一个函数，接收字符串和函数类型：f strConvert
func testFunc2(str string, f strConvert) string {
	return f(str)
}

func funcParam(str string) string {
	result := ""
	for i, val := range str {
		if i%2 == 0 {
			// 偶数位置转成大写
			result += strings.ToUpper(string(val))
		} else {
			// 奇数位置换成小写
			result += strings.ToLower(string(val))
		}
	}
	return result
}
