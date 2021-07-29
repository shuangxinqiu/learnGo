package test

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

/*
条件变量（Cond）
与互斥锁不同，条件变量的作用并不是保证在同一时刻仅有一个线程访问某一个共享数据，而是在对应的共享数据的状态发生变化时，通知其他因此而被阻塞的线程。
条件变量总是与互斥锁组合使用，互斥锁为共享数据的访问提供互斥支持，而条件变量可以就共享数据的状态的变化向相关线程发出通知。

使用场景：我需要完成一项任务，但是这项任务需要满足一定条件才可以执行，否则我就等着。

方法列表
方法名							描述
NewCond(l Locker) *Cond			生成一个cond需要传入实现Locker接口的变量。一般是*Mutex或*RWMutex类型的值。
(c *Cond) Wait()				等待通知
(c *Cond) Signal()				发生单个通知
(c *Cond) Broadcast()			广播(多个通知)
*/
func TestSyncCond(t *testing.T) {
	//声明互斥锁
	var mutex sync.Mutex
	//声明条件变量
	cond := sync.NewCond(&mutex)
	for i := 1; i <= 10; i++ {
		go func(i int) {
			//获取锁
			cond.L.Lock()
			//释放锁
			defer cond.L.Unlock()
			//等待通知，阻塞当前协程
			cond.Wait()
			//等待通知后打印输出
			fmt.Printf("输出：%d \n",i)
		}(i)
	}

	//单个通知
	time.Sleep(time.Second)
	fmt.Println("单个通知A！")
	cond.Signal()
	time.Sleep(time.Second)
	fmt.Println("单个通知B！")
	cond.Signal()

	//广播通知
	time.Sleep(time.Second)
	fmt.Println("广播通知！并睡眠1秒，等待其他协程输出！")
	cond.Broadcast()
	//等待其他协程处理完
	time.Sleep(time.Second)
	fmt.Println("运行结束！")
}


/*
一次（Once）
sync.Once是使Go方法只执行一次的对象实现，作用与init函数类似，但也有所不同。区别如下：
1.init函数是在文件包首次被加载的时候执行，且只执行一次
2.sync.Once是在代码运行中需要的时候执行，且只执行一次

方法介绍
方法名						描述
(o *Once) Do(f func())		函数只会执行一次，并保证在返回时，传入Do的函数已经执行完成。多个goroutine同时执行once.Do的时候，可以保证抢占到once.Do执行权的goroutine执行完once.Do后，其他goroutine才能的返回。
*/

func TestSyncOnce(t *testing.T) {
	echo := func() {
		t:=time.Now().Unix()
		fmt.Printf("输出时间：%v",strconv.FormatInt(t,10))
	}
	var one sync.Once
	//虽然是遍历调研，但是只会执行一次
	for i := 1; i < 10; i++ {
		go func(a,b int) {
			one.Do(echo)
		}(i,i+1)
	}

	time.Sleep(time.Second*3)
	fmt.Printf("运行结束")
}
