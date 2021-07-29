package test

import (
	"fmt"
	"testing"
)

/*
go语言同时有函数和方法，方法的本质是函数，但是方法和函数又有所不同
*/

/*
方法和函数的区别
	1.函数function是一段具有独立功能的代码，可以被反复多次调用，从而实现代码复用
	2.方法method是一个类的行为功能，只有该类的对象才能调用
	3.方法有接受者，而函数无接受者
	4.函数不可以重名，而方法可以重名。只要接受者不同，方法命就可以相同
*/

/*
基本语法
fun (接受者变量 接受者类型) 方法名(参数列表) (返回参数) {
		函数体
	}

接受者变量：接受者变量在命名时，官方建议使用接受者类型的第一个小写字母，而不是self、this之类的命名。例如：Socket类型的接受者变量可以命名为s，Connector类型的接受者变量应该命名为c等
接受者类型：接受者类型和参数类似，可以是指针类型和非指针类型
方法名、参数列表、返回参数：格式与函数一致
*/
type Student struct {
	name string
	age  int
}

//定义一个方法（接受者为Student的指针
func (s *Student) updateName(newName string) {
	s.name = newName
}

//定义一个方法（接受者为Student
func (s Student) updateAge(newAge int) {
	s.age = newAge
	fmt.Printf("修改结构体s的age -> %v \n", s)
}

/*
可以看出:  若方法的接受者不是指针，实际只是获取了一个拷贝，而不能真正改变接受者中原来的数据。
*/
func TestMethod(t *testing.T) {
	//初始化结构体
	s := Student{
		name: "leo",
		age:  22,
	}
	fmt.Printf("结构体初始化s -> %v \n", s)

	s.updateName("new leo")
	fmt.Printf("调用updateName后 -> %v \n", s)

	s.updateAge(32)
	fmt.Printf("调用updateAge后 -> %v \n", s)

}

/*
方法继承
方法是可以继承的，如果匿名字段实现了一个方法，那么包含这个匿名字段的struct也能调用该匿名字段中的方法
*/
type People struct {
	name, position string
	age            int
}
type Student2 struct {
	People
}
type Teacher struct {
	People
}

//定义一个方法
func (p People) say() {
	fmt.Printf("我叫 %s  %d岁 从事: %s \n", p.name, p.age, p.position)

}
func TestMethodExtend(t *testing.T) {
	student := Student2{People{
		name:     "xiaoqiu",
		position: "student",
		age:      11,
	}}
	teacher := Teacher{People{
		name:     "wanglaoshi",
		position: "teacher",
		age:      30,
	}}

	student.say()
	teacher.say()
}

/*
方法重写
方法重写是指一个包含了匿名字段的struct也实现了该匿名字段实现的方法（即子类也实现了父类的方法）
*/
func (s Student2) say() {
	fmt.Printf("我是一名学生,名字叫: %s 今年: %d岁 \n", s.name, s.age)

}
func TestRewriteMethod(t *testing.T) {
	student := Student2{People{"张三", "学生", 15}}
	teacher := Teacher{People{"李杨", "老师", 35}}
	// 调用方法(重写父类方法)
	student.say()
	// 调用方法(继承父类)
	teacher.say()
}
