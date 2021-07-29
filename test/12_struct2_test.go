package test

import (
	"fmt"
	"testing"
)

/*
结构体作为函数参数，若复制一份传递到函数中，在函数中对参数进行修改，不会影响到实际参数，证明结构体是值类型。
*/
func TestStructParam(t *testing.T) {
	s := student{
		name: "qiu",
		age:  10,
	}
	fmt.Printf("变量s--> 值: %v \n", s)
	grownUp(s)
	fmt.Printf("调用函数后,变量s--> 值: %v \n", s)

}

func grownUp(s student) {
	s.age = 80
	s.name = "长大的 " + s.name
}

//传结构体指针作为参数
func TestStructPointParam(t *testing.T) {
	s := student{"leo", 12, []string{}}
	fmt.Printf(" 变量s--> 值: %v \n", s)
	// 取址
	grownUpPoint(&s)
	fmt.Printf("调用函数后,变量s--> 值: %v \n", s)
}

func grownUpPoint(s *student) {
	s.age = 80
	s.name = "长大的 " + s.name
}

//返回对象
func TestStuctReturnObj(t *testing.T) {
	s := getStudent("杨过", 11, []string{"骑大雕"})
	fmt.Printf("函数返回值 s--> 值：%v 类型： %T \n", s, s)
}

func getStudent(name string, age int, likes []string) student {
	return student{name, age, likes}
}

//返回指针
func TestStructReturnPoint(t *testing.T) {
	s := getPointStudent("杨过", 40, []string{"骑大雕"})
	fmt.Printf("函数返回值 s--> 值: %v  类型: %T \n", s, s)

}

func getPointStudent(name string, age int, likes []string) *student {
	return &student{name, age, likes}
}

//匿名结构体
/*
变量名 := struct {
		//定义成员属性
	}{初始化成员属性}
*/
func TestNoNameStruct(t *testing.T) {
	s := struct {
		name, home, phone string
		age               int
	}{
		name:  "qiushuangxin",
		home:  "guangzhou",
		phone: "123123",
		age:   12,
	}
	fmt.Printf("变量 s--> 值: %v  类型: %T \n", s, s)

}

//匿名字段
/*
	匿名字段就是在结构体中的字段没有名字，只包含一个没有字段名的类型。这些字段被称为匿名字段。在同一个结构体中，同一个类型只能由一个匿名字段。
*/
type people struct {
	name, home string
	int
	float32
}

func TestStructNoNameField(t *testing.T) {
	s := people{
		name:    "zhangsan",
		home:    "guangzhou",
		int:     11,
		float32: 1212,
	}
	fmt.Printf("变量 s--> 值: %v \n", s)
	// 声明初始化匿名结构体(省略属性名)
	s2 := people{"李四", "南京", 22, 1.80}
	fmt.Printf("变量 s2--> 值: %v \n", s2)
	fmt.Printf("变量 s2--> 值: %v \n", s2.int)
}


