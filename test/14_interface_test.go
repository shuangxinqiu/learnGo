package test

import (
	"fmt"
	"testing"
)

/*
接口是一组方法签名。接口指定了类型应该具有的方法，类型决定了如何实现这些方法。当某个类型为接口中的所有方法提供了具体的实现细节时，这个类型就被称为实现了该接口。
*/

/*
使用注意事项
	1.接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量（实例）
	2.接口中所有的方法都没有方法体，即都是没有实现的方法
	3.在Go中，一个自定义类型需要将某个接口的所有方法都实现，我们说这个自定义类型实现了该接口
	4.一个自定义类型只有实现了某个接口，才能将自定义类型的实例(变量)赋给接口类型
	5.只要时自定义数据类型，就可以实现接口，不仅仅是结构体类型
	6.一个自定义类型可以实现多个接口
	7.Go接口中不能有任何变量
	8.interface类型默认是一个指针(引用类型),如果没有对interface初始化就使用，那么会输出nil
	9.空接口interface{}没有任何方法，所以所有类型都实现了空接口，即我们可以把任何一个变量赋给空接口
*/

/*
定义一个接口
*/
type Birder interface {
	fly()               //无参数无返回值方法
	eat(string2 string) //有参数无返回值方法
	//walk(string2 string) string //有参数有返回值方法
}

//定义乌鸦结构体
type Crow struct {
	name string
}

//--------------下面开始实现Birder接口----------
func (c Crow) fly() {
	fmt.Printf("我是 %s,我会飞....\n", c.name)
}
func (c Crow) eat(food string) {
	fmt.Printf("我是 %s,我喜欢吃 %s \n", c.name, food)
}

//--------------实现鸟类接口的所有方法，就代表实现了接口--------

func TestInterface(t *testing.T) {
	crow := Crow{"乌鸦"}
	crow.fly()
	crow.eat("谷子")
}

/*
模拟多态
多态：如果有几个相似而不完全相同的对象，有时人们要求在向它们发出同一个消息时，它们的反应各不相同，分别执行不同的操作，这种情况就是多态现象。Go语言中的多态性是在接口的帮助下实现的---定义接口类型，创建实现该接口的结构体对象。
*/
type Flying interface {
	getName() string
}

type bird struct {
	name string
}

func (b bird) getName() string {
	return b.name
}

type plane struct {
	name string
}

func (p plane) getName() string {
	return p.name
}

type ufo struct {
	name string
}

func (u ufo) getName() string {
	return u.name
}

func print(flyList []Flying) {
	for _, value := range flyList {
		fmt.Printf("我是%s,我会飞.....\n", value.getName())
	}
}

func TestDuoTai(t *testing.T) {
	bird := bird{name: "鸟"}
	plane := plane{name: "飞机"}
	ufo := ufo{name: "UFO"}

	flyList := make([]Flying, 0, 3)
	flyList = append(flyList, bird, plane, ufo)

	print(flyList)
}

/*
空接口是接口类型的特殊形式，空接口没有任何方法，因此任何类型都无须实现空接口。从实现的角度来看，任何值都满足这个接口的需求。因此空接口类型可以保存任何值，也可以从空接口中取出原值。
*/
//定义空接口
type Empty interface{}

//保存任意类型
func TestEmptyInterface(t *testing.T) {
	//声明变量
	var e Empty
	//保存整型
	e = 10
	fmt.Printf("保存整型： %v \n", e)
	//保存字符串
	e = "hello world"
	fmt.Printf("保存字符串： %v \n", e)
	//保存数组
	e = [...]int{1, 2, 3}
	fmt.Printf("保存数组： %v \n", e)
	//保存切片
	e = []int{11, 22, 33}
	fmt.Printf("保存切片： %v \n", e)
	//保存map
	e = map[string]int{
		"zhangsan": 11,
		"lisi":     22,
	}
	fmt.Printf("保存map： %v \n", e)
	//保存结构体
	e = struct {
		name string
		age  int
	}{"王麻子", 20}
	fmt.Printf("保存结构体: %v \n", e)

	//声明一个空接口切片
	var ee []Empty
	//保存任意类型数据到切片中
	ee = append(ee, 23, []string{"php", "go"}, map[string]int{"a": 1, "b": 2}, struct {
		city, province string
	}{"合肥", "安徽"})
	fmt.Printf("空接口切片: %v \n", ee)

}

/*
从空接口中取值，如果直接取出指定类型的值时，会发生编译错误
*/
type I interface {
}

//错误示例
func TestErrorCase(t *testing.T) {
	//声明变量num
	num := 10
	//把变量num存到空接口中
	var i I = num
	fmt.Printf("输出变量i: %v \n", i)
	//从空接口中取出值，赋值给新的变量
	//var c int = i  //!!!这里会报错
}

//正确示例
func TestCorrectCase(t *testing.T) {
	num := 10
	var i I = num
	var c int = i.(int)
	fmt.Printf("输出变量c: %v \n", c)
}

/*
接口对象转换
语法：
	方式一
	instance,ok := 接口对象.(实际类型)
	方式二
	接口对象.(实际类型)
*/
func TestInterfaceToObj(t *testing.T) {
	var a I
	// 保存整型
	a = 10
	printType(a)
	printType2(a)
	// 保存字符串
	a = "hello word"
	printType(a)
	printType2(a)
	// 保存数组
	a = [3]float32{1.0, 2.0, 3.0}
	printType(a)
	printType2(a)
	// 保存切片
	a = []string{"您", "好"}
	printType(a)
	printType2(a)
	// 保存Map
	a = map[string]string{
		"张三": "男",
		"小丽": "女",
	}
	printType(a)
	printType2(a)
	// 保存结构体
	a = people{
		name:    "leo",
		home:    "gz",
		int:     11,
		float32: 1212,
	}
	printType(a)
	printType2(a)

}

//方式一
func printType(i I) {
	if t, ok := i.(int); ok {
		echo(t)
	} else if t, ok := i.(string); ok {
		echo(t)
	} else if t, ok := i.(map[string]string); ok {
		echo(t)
	} else if t, ok := i.([]int); ok {
		echo(t)
	} else if t, ok := i.([3]string); ok {
		echo(t)
	} else if t, ok := i.(people); ok {
		echo(t)
	}
}

// 方式二
func printType2(i I) {
	switch i.(type) {
	case int:
		echo2(i)
	case string:
		echo2(i)
	case map[string]string:
		echo2(i)
	case []int:
		echo2(i)
	case [3]string:
		echo2(i)
	case people:
		echo2(i)
	}
}
func echo(i interface{}) {
	fmt.Printf("方式一 ---> 变量i类型: %T 值: %v \n", i, i)
}
func echo2(i interface{}) {
	fmt.Printf("方式二 ---> 变量i类型: %T 值: %v \n", i, i)
}
