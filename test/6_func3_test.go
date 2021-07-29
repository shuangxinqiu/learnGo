package test

import (
	"fmt"
	"testing"
)

//调用匿名函数
func TestFunc3(t *testing.T) {
	//定义一个匿名函数
	addFunc := func(a, b int) int {
		return a + b
	}

	//调用匿名函数
	add := addFunc(1, 2)
	fmt.Println("add = ", add)
}

//匿名函数当参数
func TestFunc4(t *testing.T) {
	//定义一个匿名函数，作为变量
	addFunc := func(a, b int) int {
		return a + b
	}

	//定义一个匿名函数接收匿名函数作为参数：
	//下面代码，相当于先定义匿名函数f，然后再调用匿名函数f(...)
	total := func(a, b int, f func(int, int) int) int {
		return f(a, b)
	}(10, 20, addFunc)
	fmt.Println("返回结果: ", total)
}

//递归函数
/*
递归函数必须满足以下两个条件：
在每一次调用自己时，必须是(在某种意义上)更接近于解
必须有一个终止处理或计算的准则
*/

/*
使用递归函数需要注意防止栈溢出。在计算机中，函数调用是通过栈(stack)这种数据结构实现的，每当进入一个函数调用，栈就会加一层，每当函数返回，栈就会减一层。由于栈的大小不是无限的，所以，递归调用的次数过多，会导致栈溢出。
*/

//实现阶乘
func TestFunc5(t *testing.T) {
	res := factorial(10)
	fmt.Println("10的阶乘：",res)
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	res := num * factorial(num-1)
	return res
}
