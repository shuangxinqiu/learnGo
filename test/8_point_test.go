package test

import (
	"fmt"
	"testing"
)

/*
go语言指针的最大特点是：指针不能运算
*/

func TestPoint(t *testing.T) {
	//声明一个变量
	a := 100
	fmt.Printf("变量a--> 类型: %T, 值: %v 内存地址: %p \n", a, a, &a)

	//声明一个指向整型的指针变量
	var p *int
	//取出变量a的内存地址，赋值给指针变量p
	p = &a
	fmt.Printf("指针变量p--> 类型: %T, 值: %v 指针指向的值: %v \n", p, p, *p)

	//通过指针变量修改变量a的值
	*p = 50
	fmt.Printf("通过指针编排p修改变量a的值--> 类型: %T, 值: %v 内存地址: %p \n", a, a, &a)

	var p2 *int
	if p2 == nil {
		fmt.Println("空指针")
	}
}

//使用指针作为函数的参数
func TestPointByParam(t *testing.T) {
	a := 100
	fmt.Printf("变量a--> 类型: %T, 值: %v 内存地址: %p \n", a, a, &a)
	usePointParam(&a)
	fmt.Printf("变量a--> 类型: %T, 值: %v 内存地址: %p \n", a, a, &a)

}

func usePointParam(val *int) {
	*val++
}

//指针数组
func TestPointArr(t *testing.T) {
	arrStr := [3]string{"go", "java", "php"}

	var ptArr [3]*string

	for i := 0; i < 3; i++ {
		ptArr[i] = &arrStr[i]
	}
	fmt.Printf("ptrArr类型: %T ptrArr类型值: %v \n",ptArr,ptArr)

	for i := 0; i < 3; i++ {
		a := ptArr[i]
		// 通过指针修改原数组的值
		*a += "-v"
		ptArr[i] = a
	}
	fmt.Printf("修改后: arr类型: %T arr类型值: %v \n",arrStr,arrStr)

}
