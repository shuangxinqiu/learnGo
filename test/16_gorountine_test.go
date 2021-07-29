package test

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"
	"time"
)

/*
协程又称为微线程，是一种比线程更加轻量级的存在。正如一个进程可以拥有多个线程，
协程是编译器级的，进程和线程是操作系统级的。协程不被操作系统内核管理，而完全由程序控制，因此没有线程切换的开销。

Go语言中的协程叫作Goroutine，Goroutine由Go程序运行时（runtime）调度和管理，Go程序会智能地将Goroutine中的任务合理地分配给每个CPU。创建Goroutine的成本很小，每个Goroutine
的堆栈只有几kb，且堆栈可以根据应用程序的需要增长和收缩。
*/

/*
Goroutine与Coroutine的区别
1.G能并行执行，C只能顺序执行
2.G可在多线程环境产生，C只能发生在单线程
3.C程序需要主动交出控制权，系统才能获得控制权并将控制权交给其他C
4.C的运行机制属性协作式任务处理，应用程序在不使用CPU时，需要主动交出CPU使用权。如果开发者无意间让应用程序长时间占用CPU，操作系统也无能为力，计算机很容易失去响应或死机
5.G属于抢占式任务处理，和现有的多线程和多进程任务处理非常类似。应用程序对CPU的控制最终由操作系统来管理，如果操作系统发现一个应用程序长时间占用CPU，那么用户有权终止这个任务。
*/

/*
普通函数创建goroutine
使用go关键字创建goroutine时，被调用的函数往往没有返回值，如果有返回值也会被忽略。
如果需要在goroutine中返回数据，必须使用channel，通过channel把数据从goroutine中作为返回值传出
*/
func TestGoroutine(t *testing.T) {
	go echoNum("qiu")
	go echoNum("huang")
	go echoNum("liang")
	//睡眠1秒
	time.Sleep(1 * time.Second)
	fmt.Println("运行结束")
}
func echoNum(who string) {
	for i := 1; i < 3; i++ {
		fmt.Println(who + " " + strconv.Itoa(i))
	}
}

/*
匿名函数创建goroutine
go func(){}()
*/
func TestNoNameGoroutine(t *testing.T) {
	go func(who string) {
		for i := 1; i < 3; i++ {
			fmt.Println(who + " " + strconv.Itoa(i))
		}
	}("张三")
	go func(who string) {
		for i := 1; i < 3; i++ {
			fmt.Println(who + " " + strconv.Itoa(i))
		}
	}("李四")
	go func(who string) {
		for i := 1; i < 3; i++ {
			fmt.Println(who + " " + strconv.Itoa(i))
		}
	}("王五")

	time.Sleep(1*time.Second)
	fmt.Println(runtime.NumCPU())
	fmt.Println("运行结束")
}


/*
并发运行性能调整
设置运行的cpu数：为了充分利用多cpu的优势，在golang程序中，可以通过runtime.GOMAXPROCS()函数设置运行的cpu数目

runtime.GOMAXPROCS( 逻辑CPU)
逻辑CPU可以通过 runtime.NumCpu()函数获取

数值	含义	示例
< 1	不修改任何数值	runtime.GOMAXPROCS(0)
= 1	单核执行	runtime.GOMAXPROCS(1)
> 1	多核并发执行	runtime.GOMAXPROCS(4) 4核并发执行

Go 1.5版本之前，默认使用的是单核心执行。从Go 1.5版本开始，默认执行runtime.GOMAXPROCS（逻辑CPU数量）,让代码并发执行，最大效率地利用CPU。
GOMAXPROCS同时也是一个环境变量，在应用程序启动前设置环境变量也可以起到相同的作用。
*/
/*
注意事项：
1.所有goroutine在main()函数结束时会一同结束
2.如果需要在goroutine中返回数据，必须使用channel，通过channel把数据从goroutine中作为返回值传出
3.启动多个goroutine时，会随机调度
*/
