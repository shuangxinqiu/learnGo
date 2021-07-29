package test

import (
	"fmt"
	"strconv"
	"testing"
)

/*
select语句的机制有点像switch语句，不同的是，select会随机挑选一个可通信的case来执行，如果所有case都没有数据到达，则执行default，如果没有default语句，select就会阻塞，直到有case接收到数据。
*/

//select使用，没有default
//如果没有defau语句时，select就会阻塞，直到有case接收到数据
func TestSelect(t *testing.T) {
	//创建通道
	intChan := make(chan int)
	stringChan := make(chan string)

	//写入
	go func() {
		for i := 1; i < 4; i++ {
			intChan <- i
			stringChan <- "邱xx-" + strconv.Itoa(i)
		}
	}()

	//使用select 没有default会阻塞等待。（会随机运行一次）
	select {
	case n := <-intChan:
		fmt.Printf("接收到intChan中的数据：%v \n", n)
	case s := <-stringChan:
		fmt.Printf("接收到stringChan中的数据：%v \n", s)
	}
	fmt.Printf("运行结束！")
}

/*
select使用，有default
如果default时，所有case都没有数据到达，则执行default
*/
func TestSelectChannel(t *testing.T) {
	//创建通道
	intChan := make(chan int)
	stringChan := make(chan string)

	//写入
	go func() {
		for i := 1; i < 4; i++ {
			intChan <- i
			stringChan <- "邱xxx-"+strconv.Itoa(i)
		}
	}()

	// 使用select 没有default,则会阻塞等待。（会随机运行一次）
	select {
	case n := <-intChan:
		fmt.Printf("接收到intChan中的数据: %v\n", n)
	case s := <-stringChan:
		fmt.Printf("接收到stringChan中的数据: %v\n", s)
	default:
		fmt.Println("什么数据都没收到!")
	}
	fmt.Printf("运行结束!")
}