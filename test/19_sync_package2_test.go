package test

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

/*
读写锁（RWMutex）
RWMutex是读写互斥锁，简称读写锁。该锁可以同时被多个读取者或被唯一个写入者持有。RWMutex类型锁跟Goroutine无关，可以由不同的Goroutine加锁、解锁。RWMutex也可以创建为其他结构体的字段；零值为解锁状态。

RWMutex锁结构
type RWMutex struct{
	w Mutex //用于控制多个写锁，获得写锁首先要获取该锁，如果有一个写锁在进行，那么再来到的写锁将会阻塞于此
	writerSem uint32 //写阻塞等待的信号量，最后一个读者释放锁时会释放信号量
	readerSem uint32 //读阻塞的协程等待的信号量，持有写锁的协程释放锁后会释放信号量
	readerCount int32 //记录读者个数
	readerWait int32 //记录写阻塞时读者个数
}

读写锁堵塞场景
1.写锁需要阻塞写锁：一个协程拥有写锁时，其他协程写锁需要阻塞
2.写锁需要阻塞读锁：一个协程拥有写锁时，其他协程读锁需要阻塞
3.读锁需要阻塞写锁：一个协程拥有读锁时，其他协程写锁需要阻塞
4.读锁不能阻塞读锁：一个协程拥有读锁时，其他协程也可以拥有读锁

方法列表
方法名							描述
(rw *RWMutex) RLock()			获取读锁，当一个协程拥有读锁时，其他协程写锁需要阻塞
(rw *RWMutex) RUnlock()			释放读锁
(rw *RWMutex) Lock()			获取写锁，与Mutex完全一致；当一个协程拥有写锁时，其他协程读写锁都需要阻塞
(rw *RWMutex) Unlock()			释放写锁
*/
//不作为结构体属性使用

//声明全局变量，文件内容
var fileContext string

//声明全局读写互斥锁
var rwMutex sync.RWMutex

//声明全局等待组,已存在
func TestRWMutex(t *testing.T) {
	//设置计数器
	wg.Add(5)
	for i := 1; i <= 5; i++ {
		name := "同学-" + strconv.Itoa(i)
		if i%2 == 0 {
			go readFile(name)
		} else {
			go writeFile(name, strconv.Itoa(i))
		}
	}
	wg.Wait()
	fmt.Printf("运行结束！！！！")
}

func readFile(name string) {
	//释放读锁
	defer rwMutex.RUnlock()
	//获取读锁
	rwMutex.RLock()
	//打印读取内容
	fmt.Printf("%s 获取读锁，读取内容为：%s \n", name, fileContext)
	//计数器减一
	wg.Done()
}
func writeFile(name, s string) {
	//释放写锁
	defer rwMutex.Unlock()
	//获取读锁
	rwMutex.Lock()
	//写入内容
	fileContext = fileContext + " " + s
	fmt.Printf("%s 获取写锁，写入内容: %s。 文件内容变成: %s \n", name, s, fileContext)
	wg.Done()
}

//作为结构体属性使用
//定义一个文件结构体
type fileResource struct {
	content string
	wg      sync.WaitGroup
	rwLock  sync.RWMutex
}

//读文件
func (file *fileResource) read(name string) {
	//释放读锁
	defer file.rwLock.RUnlock()
	//获取读锁
	file.rwLock.RLock()
	//打印读取内容
	time.Sleep(time.Second)
	fmt.Printf("%s 获取读锁，读取内容：%s \n", name, file.content)
	//计数器减一
	file.wg.Done()
}

//写文件
func (file *fileResource) write(name, s string) {
	//释放写锁
	defer file.rwLock.Unlock()
	//获取写锁
	file.rwLock.Lock()
	//写入内容
	//time.Sleep(time.Second)
	file.content = file.content + " " + s
	fmt.Printf("%s 获取写锁，写入内容：%s。 文件内容变成：%s \n", name, s, file.content)
	//计数器减一
	file.wg.Done()
}
func TestUseRWMutexByStruct(t *testing.T) {
	//声明结构体
	var file fileResource
	//设置计数器
	file.wg.Add(5)
	for i := 1; i <= 5; i++ {
		name := "同学-" + strconv.Itoa(i)
		if i%2 == 0 {
			go file.read(name)
		} else {
			go file.write(name, strconv.Itoa(i))
		}
	}
	file.wg.Wait()
	fmt.Printf("运行结束了！！")
}
