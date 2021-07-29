package test

import (
	"fmt"
	"testing"
	"time"
)

/*
Channel的阻塞特性
1.channel默认是阻塞的
2.当数据被发送到channel时会发送阻塞，直到有其他Goroutine从该channel中读取数据
3.当从channel读取数据时，读取也会被阻塞，直到其他Goroutine将数据写入该channel
*/
func TestChannel2(t *testing.T) {
	intChan := make(chan int)
	bookChan := make(chan bool)

	go func(intChan chan int) {
		intChan <- 50
		fmt.Printf("写入数据：50 \n")
		close(intChan)
		fmt.Println("写协程结束")
	}(intChan)

	go func(intChan chan int, boolChan chan bool) {
		data, ok := <-intChan
		if ok {
			fmt.Println("读取数据：", data)
			bookChan <- true
			close(bookChan)
			fmt.Println("读协程结束")
		}
	}(intChan, bookChan)

	<-bookChan // 忽略接收，达到阻塞的效果。(如果不阻塞，则会直接输出: 程序运行结束!,不会等待协程执行)
	//阻塞channel等待匿名函数的Goroutine运行结束，防止主函数的Goroutine退出而导致匿名函数的Goroutine提前退出。
	fmt.Println("主线程运行结束")
}

/*
关闭channel
发送方写入完毕后需要主动关闭channel，用于通知接收方数据传递完毕。接收方通过data,ok := <- ch判断channel是否关闭，如果ok=false，则表示channel已经被关闭。

示例可以看出: 可以从关闭后的channel中继续读取数据，取到的值为该类型的零值。比如整型是:0; 字符串是：""

注：
1.往关闭的channel中写入数据会报错：panic: send on closed channel。导致程序崩溃。
2.重复关闭chan，会崩溃
*/
func TestCloseChannel(t *testing.T) {
	// 创建一个整数型chan
	intChan := make(chan int)
	// 创建一个写入channel的协程
	go func(intChan chan int) {
		intChan <- 10
		fmt.Printf("发送数据: 10 \n")
		intChan <- 20
		fmt.Printf("发送数据: 20 \n")
		// 关闭通道
		close(intChan)
	}(intChan)

	// 读取数据
	a := <-intChan
	fmt.Printf("接收数据: %v \n", a)
	b := <-intChan
	fmt.Printf("接收数据: %v \n", b)

	// 此时的Chan已经关闭，而且里面的数据也都已经取完
	c := <-intChan
	fmt.Printf("接收数据: %v \n", c)
	fmt.Println("程序运行结束!")
}

/*
缓冲channel
默认创建的都是非缓冲channel，读写都是即时阻塞。缓冲channel自带一块缓冲区，可以暂时存储数据，如果缓冲区满了，就会发送阻塞。
缓冲通道在发送时无须等待接收方接收即可完成发送过程，并且不会发生阻塞，只有当缓冲区满时才会发生阻塞。同理，如果缓冲通道中有数据，接收时将不会发生阻塞，直到通道中没有数据可读时，通道将会再度阻塞。

语法：cha1 := make(chan T,n) n代表缓冲区大小
*/
func TestBufferChannel(t *testing.T) {
	fmt.Printf("开始时间：%v \n", time.Now().Unix())
	//创建一个缓冲区为2的整型chan
	intChan := make(chan int, 2)
	//不会发生阻塞，因为缓冲区未满
	intChan <- 100
	fmt.Printf("结束时间：%v \n", time.Now().Unix())
	fmt.Printf("intChan 类型：%T 缓冲大小：%v \n", intChan, cap(intChan))
	fmt.Printf("程序运行结束")
}

/*
单向channel
channel默认是双向的，即可读可写。定向channel也叫单向channel，只读或只写。直接创建单向channel没有任何意义。通常的做法是创建双向channel，然后以单向channel的方式进行函数传递

只读： ch <- chan int
只写：ch chan <- int
*/
func TestSingleChannel(t *testing.T) {
	//创建一个整数型chan
	intChan := make(chan int)
	go writeChan(intChan)
	go readChan(intChan)
	time.Sleep(50 * time.Second)
	fmt.Printf("运行结束")
}

//定义只读通道的函数
func readChan(ch <-chan int) {
	for data := range ch {
		fmt.Printf("读出数据：%v \n", data)
	}
}

//定义只写通道的函数
func writeChan(ch chan<- int) {
	for i := 1; i < 50; i++ {
		ch <- i
		fmt.Printf("写入数据：%v \n", i)
	}
}

/*
计时器与channel
计时器类型表示单个事件。当计时器过期时，当前时间将被发送到c上（c是一个只读channel <- chan time.Time，该channel中放入的是Timer结构体），除非计时器是After(
)创建的。计时器必须使用NewTimer()或After()创建。
*/
//NewTimer()创建一个新的计时器，它会在至少持续时间d之后将当前时间发送到其channel上。
func TestNewTimerChannel(t *testing.T) {
	//创建一个计时器
	timer := time.NewTimer(5 * time.Second)
	fmt.Printf("开始时间：%v \n", time.Now())
	//此处会阻塞5秒
	out := <-timer.C
	fmt.Printf("变量out-> 类型：%T 值：%v \n", out, out)
	fmt.Printf("结束时间：%v \n", time.Now())
}

/*
After()函数相当于NewTimer(d).C，如下源码
func After(d Duration) <- chan Time {
return NewTimer(d).C
}
*/
func TestAfterChannel(t *testing.T) {
	//创建一个计时器，返回的是chan
	ch := time.After(5 * time.Second)
	fmt.Printf("开始时间：%v \n", time.Now())
	//此处会阻塞5秒
	out := <-ch
	fmt.Printf("变量out -> 类型：%T 值：%v \n", out, out)
	fmt.Printf("结束时间: %v \n", time.Now())
}
