package test

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

/*
sync包提供了互斥锁。除了Once和Wait和WaitGroup类型，其余多数使用于低水平的程序，多数情况下，高水平的同步使用channel通信性能会更优一些
*/

/*
WaitGroup（并发等待组），即等待一组Goroutine结束。父Goroutine调用Add()方法来设置应等待Goroutine的数量。每个被等待的Goroutine在结束时应该调用Done()方法。与此同时，主Goroutine可调用Wait()方法阻塞至所有Goroutine结束。

type WaitGroup struct {
	noCopy noCopy
	state1 [3]uint32
}

方法列表
方法名									功能
(wg *WaitGroup) Add(delta int)			等待组的计数器+1
(wg *WaitGroup) Done()					等待组的计数器-1
(wg *WaitGroup) Wait()					当等待组计数器不等于0时，阻塞直到0

Add参数取值范围
等待组内部拥有一个计数器，计数器的值可以通过Add(delta int)方法调用实现计数器的增加和减少。该方法应该在创建新的Goroutine之前调用。

参数值x取值
取值		描述
delta<0		x小于0，但会报错：panic:sync:negative WaitGroup counter
delta=0 	x等于0，会释放Wait()方法阻塞等待的所有Goroutine
delta>0 	x大于0，Wait()方法会阻塞Goroutine直到WaitGroup计数减为0
*/
func TestNoUseWaitGroup(t *testing.T) {
	//创建通道
	intChan := make(chan int)

	go func(intChan chan int) {
		sum := 0
		for i := 1; i <= 50; i++ {
			sum += i
		}
		intChan <- sum
	}(intChan)

	go func(intChan chan int) {
		sum := 0
		for i := 51; i <= 100; i++ {
			sum += i
		}
		intChan <- sum
	}(intChan)

	go func(intChan chan int) {
		sum1 := <-intChan
		sum2 := <-intChan
		fmt.Printf("sum1 = %d sum2 = %d  \nsum1 + sum2 = %d \n", sum1, sum2, sum1+sum2)
	}(intChan)

	//注意，需要手动阻塞sleep
	time.Sleep(time.Second)
	fmt.Printf("运行结束！")
}

func TestUseWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)

	intChan := make(chan int)
	go func(intChan chan int, wg *sync.WaitGroup) {
		sum := 0
		for i := 1; i <= 50; i++ {
			sum += i
		}
		intChan <- sum
		wg.Done()
	}(intChan, &wg)

	go func(intChan chan int, wg *sync.WaitGroup) {
		sum := 0
		for i := 51; i <= 100; i++ {
			sum += i
		}
		intChan <- sum
		wg.Done()
	}(intChan, &wg)

	go func(intChan chan int, wg *sync.WaitGroup) {
		sum1 := <-intChan
		sum2 := <-intChan
		fmt.Printf("sum1 = %d sum2 = %d  \nsum1 + sum2 = %d \n", sum1, sum2, sum1+sum2)

		wg.Done()
	}(intChan, &wg)

	wg.Wait()
	fmt.Printf("运行结束！！！")

}

/*
互斥锁（Mutex）
Mutex是一个互斥锁，保证同时只有一个Goroutine可以访问共享资源。Mutex类型的锁和Goroutine无关，可以由不同的Goroutine加锁和解锁。也可以为其他结构体的字段，零值为解锁状态。

type Mutex struct{
	state int32 //state 表示当前互斥锁的状态
	sema uint32 //sema 时用于控制锁状态的信号量
}

方法列表
方法名					描述
(m *Mutex) Lock()		方法锁住m，如果m已经加锁，则阻塞直到m解锁
(m *Mutex) Unlock()		解锁m，如果m未加锁会导致运行时错误
*/
/*
使用（售票）

需求：模拟多个窗口售票
*/
//不作为结构体属性使用
var mutex sync.Mutex
var wg sync.WaitGroup

var ticket int = 10

func TestUseMutex(t *testing.T) {
	wg.Add(3)
	go saleTicket("窗口A", &wg)
	go saleTicket("窗口B", &wg)
	go saleTicket("窗口C", &wg)
	wg.Wait()
	fmt.Printf("运行结束！")
}

func saleTicket(windowName string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		mutex.Lock()
		if ticket > 0 {
			time.Sleep(time.Second)
			ticket--
			fmt.Printf("%s 卖出一张票，余票: %d \n", windowName, ticket)
		} else {
			fmt.Printf("%s 票已卖完！ \n", windowName)
			mutex.Unlock()
			break
		}
		mutex.Unlock()
	}
}

//作为结构体属性使用
type ticketPool struct {
	over  int
	wg    sync.WaitGroup
	mutex sync.Mutex
}

func (t *ticketPool) saleTicket(windowName string) {
	defer t.wg.Done()
	for {
		t.mutex.Lock()
		if t.over > 0 {
			time.Sleep(time.Second)
			t.over--
			fmt.Printf("%s 卖出一张票，余票: %d \n", windowName, t.over)
		}else{
			t.mutex.Unlock()
			fmt.Printf("%s 票已卖完! \n", windowName)
			break
		}
		t.mutex.Unlock()
	}
}

func TestUseMutexByStruct(t *testing.T) {
	pool := ticketPool{
		over:  10,
		wg:    sync.WaitGroup{},
		mutex: sync.Mutex{},
	}
	windowNum := 3
	pool.wg.Add(windowNum)
	for i := 1; i <= windowNum; i++ {
		go pool.saleTicket("窗口"+strconv.Itoa(i))
	}
	pool.wg.Wait()
}
