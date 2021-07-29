package test

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

/*
错误是指程序中出现不正常的情况，从而导致程序无法正常运行。Go语言中没有try...catch来捕获错误，而是通过defer+recover+panic模式来实现捕捉错误信息。
*/

/*
error接口
Go语言通过内置的错误类型提供了非常简单的错误处理机制，即error接口。该接口的定义如下：
type error interface {
	Error() string
}

对大多数函数，如果要返回错误，大致上都可以定义为如下模式：必须将error作为多种返回值的最后一个
func Demo(参数列表...)(x T, err error) {
函数体
}

错误判断
x,err := Demo(参数列表...)
if err != nil {
打印错误信息
}
*/

/*
创建error对象几种方式：
	Go语言errors包下的New()函数可以返回error对象
	使用fmt包下的Errorof()函数
*/
func TestError(t *testing.T) {
	//方式一：使用errors包下的New()
	err := createError(1)
	printError(err)

	err2 := createError(2)
	printError(err2)

}
func createError(way int) error {
	if way == 1 {
		//方式一：使用errors包下的New()
		return errors.New("方式一：使用errors包下的New(...)")
	} else if way == 2 {
		//方式二：使用fmt包下的Errorof()
		return fmt.Errorf("方式二：使用fmt包下的Errorof(...) -->")
	}
	return nil
}
func printError(err error) {
	if err != nil {
		fmt.Printf("err==> %v | err.Error() ==> %v | 类型==> %T \n", err, err.Error(), err)
	}
}

/*
自定义错误的实现步骤
1.定义一个结构体，表示自定义错误的类型。
2.让自定义错误类型实现error接口：Error() string
3.定义一个返回error的函数，根据程序实际功能而定
*/
//第一步：定义一个结构体
type MyError struct {
	msg string
	t   time.Time
}

//第二步：结构体实现error接口
func (m MyError) Error() string {
	return fmt.Sprintf("错误信息：%s 发生时间：%v", m.msg, m.t)
}

//第三步：定义一个返回error的函数
func login(phone, pwd string) (bool, error) {
	if phone == "136" && pwd == "123123" {
		return true, nil
	}
	err := MyError{
		msg: "账号密码错误",
		t:   time.Now(),
	}
	return false, err
}
func TestMyError(t *testing.T) {
	res, err := login("136", "123123")
	if err != nil {
		fmt.Printf("登录失败--> %v  T--> %T \n", err.Error(), err)
	} else {
		fmt.Printf("登录成功--> %v  \n", res)
	}

	res2, err2 := login("1360", "123123")
	if err2 != nil {
		fmt.Printf("登录失败--> %v  T--> %T \n", err2.Error(), err2)
	} else {
		fmt.Printf("登录成功--> %v  \n", res2)
	}
}

/*
延迟函数（defer）
关键字defer用于延迟一个函数或者方法（或者当前所创建的匿名函数）的执行。
defer语句只能出现再函数或方法的内部。
多个defer语句时，当函数执行到最后时或者报错的时候，按照逆序执行，最后该函数返回

注：defer语句经常被用于处理成对的操作，如打开-关闭、连接-断开连接、加索-释放锁。特别是执行打开资源的操作时，遇到错误需要提前返回，在返回前需要关闭相应的资源，不然很容易造成资源泄露问题
*/
func TestDefer(t *testing.T) {
	defer defer1()
	defer defer2(1, 2)
	defer defer2(3, 4)

	//匿名函数
	defer func() {
		fmt.Printf("匿名函数defer3....\n")
	}()

	//正常处理代码
	for i := 1; i < 4; i++ {
		fmt.Println(i)
	}
}
func defer1() {
	fmt.Printf("函数defer1....\n")
}
func defer2(a, b int) {
	fmt.Printf("函数defer2....a=%d b=%d a+b=%d \n", a, b, a+b)
}

/*
panic（崩溃）
panic让当前的程序进入恐慌，中断程序的执行。
panic()是一个内置函数，可以中断原有的控制流程。其功能类似PHP中的throw
*/
//造成panic的场景：数组访问越界
func TestPanicArrOverSide(t *testing.T) {
	//测试访问数组下标越界
	//arr := [...]int{1,2,3}
	//数组下标最大为2
	//fmt.Println(arr[3])
}

//造成panic的场景：访问未初始化的指针或nil指针
func TestPanicPoint(t *testing.T) {
	//定义一个指针类型（默认值是nil）
	var b *int
	fmt.Println(b)
	//访问nil的指针
	fmt.Println(*b)
}

//造成panic的场景：向已经close的chan(管道)里发生数据
func TestPanicChannel(t *testing.T) {
	//测试 往已经close的chan（管道）里发生数据
	//1.声明一个channel
	var ch = make(chan int, 1)
	//2.关闭channel
	close(ch)
	//3.向已经关闭的channel写数据
	ch <- 1
}

//造成panic的场景：类型断言
func TestPanicType(t *testing.T) {
	//测试 类型断言
	var i interface{} = "hello"
	a := i.([]string)
	fmt.Println(a)
}

//使用panic
func TestUsePanic(t *testing.T) {
	throw("请求成功", 1)
	throw("请求失败", 0)

}
func throw(msg string, code int) {
	if code == 0 {
		panic("error msg: " + msg)
	}
	fmt.Println("正常输出: " + msg)
}

/*
recover(恢复)
Go语言中没有try...catch来捕获错误，一旦触发panic就会导致程序崩溃。在Go中是通过recover让程序恢复。
值得注意的是recover()必须在延迟函数defer中有效

在正常程序运行中，调用recover()会返回nil，并且没有其他任何效果。如果当前的Goroutine陷入恐慌，调用recover()可以捕获panic()的输入值，使程序恢复正常运行。
*/
func TestRecover(t *testing.T) {
	defer func() {
		err := recover()
		msg := fmt.Sprintf("err信息: %v",err)
		if err != nil {
			//程序触发panic时，会被这里捕获
			fmt.Println(msg)
		}
		fmt.Println("这里会输出")
	}()

	testPanic("请求失败")
	fmt.Println("这里不会输出")
}
func testPanic(err string) {
	panic("错误信息：" + err)
}
