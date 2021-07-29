package test

import (
	"fmt"
	"testing"
)

/*
Go没有沿袭传统面向对象编程中的诸多概念，也没有提供类(class)，但是它提供了结构体(struct)，方法(method)可以在结构体上添加。与类相似，结构体提供了捆绑数据和方法的行为。
*/

/*
结构体是由一系列相同类型或不同类型的数据构成的数据集合。
结构体的定义只是一种内存布局的描述，只有当结构体实例化时，才会真正地分配内存。因此必须在定义结构体并实例化后才能使用结构体的字段。
*/
/*
1.类型名是标识结构体的名称，在同一个包内不能重复
2.结构体的属性，也叫字段（field），必须唯一
3.同类型的成员属性可以写在一行
4.结构体是值类型
5.只有当结构体实例化时，才能使用结构体的字段
*/
type student struct {
	name string
	age  int
	like []string
}

func TestStruct(t *testing.T) {
	//使用var声明
	var s student
	fmt.Printf("变量s--> 类型： %T 值：%v \n", s, s)
	//给属性赋值
	s.name = "张三"
	s.age = 22
	s.like = []string{"games", "sport"}
	fmt.Printf("给属性赋值，变量s--> 类型：%T 值：%v \n", s, s)

	//简短声明,先实例化结构体，后赋值
	s2 := student{}
	s2.name = "李四"
	s2.age = 44
	s2.like = []string{"asd", "qweq"}
	fmt.Printf("给属性赋值，变量s2--> 类型：%T 值：%v \n", s2, s2)

	//声明时初始化
	s3 := student{
		name: "wangwu",
		age:  55,
		like: []string{"1", "2"},
	}
	fmt.Printf("给属性赋值，变量s3--> 类型：%T 值：%v \n", s3, s3)

	//声明时初始化，省略属性
	s4 := student{"qiu", 43, []string{"q", "w"}}
	fmt.Printf("给属性赋值，变量s4--> 类型：%T 值：%v \n", s4, s4)

	//使用new
	//使用内置函数new()对结构体进行实例化，结构体实例化后形成指针类型的结构体，new()内置函数会分配内存。第一个参数是类型，而不是值，返回的值是指向该类型新分配的零值的指针。
	ss := new(student)
	fmt.Printf(" 变量ss--> 类型: %T 值: %v \n",ss,ss)

	(*ss).name = "包青天"
	(*ss).age = 55
	(*ss).like = []string{"判案"}
	fmt.Printf(" 变量ss--> 类型: %T 值: %v \n",ss,ss)
	// 语法糖写法(省略*)
	ss.name = "包大人"
	ss.age = 99
	ss.like = []string{"判案","元芳你怎么看"}
	fmt.Printf(" 变量ss--> 类型: %T 值: %v \n",ss,ss)
}

