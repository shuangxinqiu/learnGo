package test

import (
	"fmt"
	"testing"
)

/*
	结构体嵌套
	将一个结构体作为另一个结构体属性（字段），这种结构体就是结构体嵌套
	结构体嵌套可以模拟面向对象编程中的以下两种关系
		聚合关系：一个类作为另一个类的属性
		继承关系：一个类作为另一个类的子类。子类和父类的关系。
*/

/*
聚合场景
模拟聚合关系时一个要采用有名字的结构体作为字段
*/
type student1 struct {
	name string
	height float32
	schoolInfo school
}

type school struct {
	schoolName, schoolAddress string
}
func TestStructMore(t *testing.T) {
	//简短声明嵌套结构体
	s := student1{
		name:       "xiaoqiu",
		height:     168,
		schoolInfo: school{"广州大学","广州"},
	}
	fmt.Printf("变量 s--> 值: %v 类型: %T \n", s,s)

	//使用var
	var ss student1
	ss.name = "xiaolong"
	ss.height =188
	ss.schoolInfo = school{
		schoolName:    "北京大学",
		schoolAddress: "北京",
	}
	fmt.Printf("变量 ss--> 值: %v 类型: %T \n", ss,ss)

	//使用new
	s2 := new(student1)
	s2.name = "小虎"
	s2.height = 1.77
	s2.schoolInfo.schoolName = "武汉大学"
	s2.schoolInfo.schoolAddress = "武汉"
	fmt.Printf("变量 s2--> 值: %v 类型: %T \n", s2,s2)
}

/*
模拟继承
在结构体中，属于匿名结构体的字段称为提升字段，它们可以被访问，匿名结构体就像是该结构体的父类
*/
type student2 struct {
	people
	class string
}

func TestStructExtend(t *testing.T) {
	// 方式1.使用new声明结构体
	var s = new(student2)
	// 集成父类成员
	s.name = "qiu"
	s.home = "gz"
	s.int =12
	s.float32 = 1212
	// 自己成员
	s.class = "grage 3"
	fmt.Printf("变量s -> %v \n",s)

	// 方式2.使用简短声明
	s2 := student2{people{"李四","13",1,2},"四年级"}
	fmt.Printf("变量s2 -> %v \n",s2)
}

//成员冲突
type A struct {
	name string
	age  int
}
type B struct {
	name   string
	height float32
}
// 在C结构体中嵌套A和B
type C struct {
	A
	B
}
func TestStruct4(t *testing.T) {
	// 定义结构体C
	c := C{}
	fmt.Printf("变量c -> %v \n", c)

	// 不冲突的成员赋值
	c.age = 12
	c.height = 1.88
	// 冲突的成员赋值
	c.A.name = "这是A的成员"
	c.B.name = "这是B的成员"
	fmt.Printf("变量c -> %v \n", c)
}