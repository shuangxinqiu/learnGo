package test

import (
	"fmt"
	"testing"
)

/*
channel即Go的通道，是协程之间的通信机制。一个channel是一条通信管道，它可以让一个协程通过它给另一个协程发送数据。每个channel都需要指定数据类型，即channel可发送数据的类型。
Go语言主张通过数据传递来实现共享内存，而不是通过共享内存来实现数据传递。
*/
/*
channel是引用类型，需要使用make()进行创建
方式1：
	var chan1 chan 数据类型
	chan1 = make(chane 数据类型)

方式2：
	chan2 := make(chan 数据类型)
*/
type ChannelPeople struct {
}

func TestChannel(t *testing.T) {
	//创建一个整数型chan
	intChan := make(chan int)
	fmt.Printf("intChan类型: %T 值: %v \n", intChan, intChan)
	//创建一个空接口chan，可以存放任意类型数据
	interfaceChan := make(chan interface{})
	fmt.Printf("interfaceChan类型：%T 值：%v \n", interfaceChan, interfaceChan)
	//创建一个指针chan
	peopleChan := make(chan *ChannelPeople)
	fmt.Printf("peopleChan类型：%T 值：%v \n", peopleChan, peopleChan)
}

/*
向channel发送数据
通过channel发送数据需要使用特殊的操作符<-,需要注意的是：channel发送的值的类型必须与channel的元素类型一致

channel变量名 <- 值
*/
func TestErrorSendCase(t *testing.T) {
	intChan := make(chan int)
	intChan <- 5
	fmt.Printf("intChan类型: %T 值: %v \n", intChan, intChan)
}

//上面实例运行会死锁，原因：如果Goroutine在一个channel上发送数据，其他的Goroutine应该接收得到数据；如果没有接收，那么程序经在运行时出现死锁

func TestCorrectSendCase(t *testing.T) {
	intChan := make(chan int)
	//写入数据（协程写入）
	go sendMsg(intChan)
	//接收数据（主线程读取）
	a := <-intChan
	fmt.Printf("接收数据：%v \n", a)
	fmt.Println("运行结束")
}
func sendMsg(msg chan int) {
	msg <- 5
	fmt.Println("写入数据：5")
}

/*
从channel接收数据
阻塞接收语法：执行该语句时channel将会阻塞，直到接收到数据并赋值给data变量
	方式一：chan 指的是通道变量
		data := <- chan
	方式二：data 表示接收到的数据。未接收到数据时，data为channel类型的零值，ok(布尔类型)表示是否接收到数据
		data, ok := <- chan


忽略接收语法：执行该语句时channel将会阻塞。其目的不在于接收channel中数据，而是为了阻塞Goroutine
<- chan

注：如果Goroutine正在等待从channel接收数据，而其他Goroutine并没有写入数据时程序将会死锁
*/

/*
循环接收
循环接收数据，需要配合使用关闭channel，借助普通for循环和for...range语句循环接收多个元素。遍历channel，遍历的结果就是接收到的数据，数据类型就是channel的数据类型。

普通for循环接收channel数据，需要有break循环的条件；
for...range会自动判断出channel已关闭，而无须通过判断来终止循环
*/

/*
使用普通for接收
1：data := <- ch
2: data,ok := <- ch
*/
func TestUseForChannel(t *testing.T) {
	//创建一个整型chan
	intChan := make(chan int)
	//写入数据,协程写入
	go func(cha chan int) {
		//写入
		for i := 1; i < 5; i++ {
			intChan <- i
			fmt.Printf("写入数据 ->  %v \n", i)

		}

		//关闭通道
		close(intChan)
	}(intChan)

	//方式一：data := <- ch
	for {
		//接收数据
		out := <-intChan
		//判断通道是否关闭
		//如果通道关闭，则out为通道类型的零值，这里是int类型，所以是0
		if out == 0 {
			fmt.Println("通道已关闭")
			break
		}
		fmt.Printf("接收数据 ==> %v \n", out)
	}
	fmt.Println("程序运行结束！")
}

func TestUseForChannel2(t *testing.T) {
	//创建一个整型chan
	intChan := make(chan int)
	//写入数据,协程写入
	go func(cha chan int) {
		//写入
		for i := 1; i < 6; i++ {
			intChan <- i
			fmt.Printf("写入数据 ->  %v \n", i)

		}

		//关闭通道
		close(intChan)
	}(intChan)

	//方式二：data,ok := <- ch
	for {
		//接收数据
		out,ok := <-intChan
		//判断通道是否关闭
		//如果通道关闭，则out为通道类型的零值，这里是int类型，所以是0
		if !ok {
			fmt.Println("通道已关闭")
			break
		}
		//time.Sleep(time.Second*1)
		fmt.Printf("接收数据 ==> %v \n", out)
	}
	fmt.Println("程序运行结束！")
}

/*
使用for...range接收
*/
func TestUseForRangeChannel(t *testing.T) {
	// 创建一个整数型chan
	intChan := make(chan int)
	// 写入数据(协程写入)
	go func(cha chan int) {
		// 写入
		for i := 1; i < 5; i++ {
			intChan <- i
			fmt.Printf("写入数据 ->  %v \n", i)
		}
		// 关闭通道
		close(intChan)
	}(intChan)
	// 使用 for...range接收
	for data := range intChan {
		fmt.Printf("接收数据 ==> %v \n", data)
	}
	fmt.Println("程序运行结束!")
}
