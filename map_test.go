package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

//工序1采购
func buy(n int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 1; i < n; i++ {
			//fmt.Println("工序1", i)
			out <- fmt.Sprint("配件", i)
		}
	}()

	return out
}

//工序2组装
func build(in <-chan string,n int) <-chan string {
	out := make(chan string)
	i := 1
	go func() {
		defer close(out)
		for c := range in {
			fmt.Println("工序2-"+strconv.Itoa(n), i)
			i++
			out <- "组装（" + c + ")"
		}
	}()
	return out
}

//工序3打包
func pack(in <-chan string) <-chan string {
	out := make(chan string)
	//i := 1
	go func() {
		defer close(out)
		for c := range in {
			//fmt.Println("工序3", i)
			//i++
			out <- "打包(" + c + ")"
		}
	}()
	return out
}

//扇入函数（组件），把多个channel中的数据发送到一个channel中
func merge(ins ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	wg.Add(len(ins))
	out := make(chan string)
	//把一个channel中的数据发送到out中
	p := func(in <-chan string,key int) {
		defer wg.Done()
		for c := range in {
			out <- "merge"+strconv.Itoa(key)+"(" + c + ")"
		}
	}
	//扇入，需要启动多个goroutine用于处于多个channel中的数据
	for key, cs := range ins {
		go p(cs,key)
	}
	//等待所有输入的数据ins处理完，再关闭输出out
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func TestMap2(t *testing.T) {
	coms := buy(10)
	phones1 := build(coms,1)
	phones2 := build(coms,2)
	phones3 := build(coms,3)
	phones := merge(phones1,phones2,phones3)
	packs := pack(phones)
	for p := range packs {
		fmt.Println(p)
	}

	time.Sleep(10*time.Second)
}

func TestMap(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(4)
	ctx, stop := context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		watchDog(ctx, "[监控狗1]")
	}()
	go func() {
		defer wg.Done()
		watchDog(ctx, "[监控狗2]")
	}()
	go func() {
		defer wg.Done()
		watchDog(ctx, "[监控狗3]")
	}()

	valCtx := context.WithValue(ctx, "UserId", 2)
	go func() {
		defer wg.Done()
		getUser(valCtx)
	}()

	time.Sleep(5 * time.Second) //先让监控狗监控5秒
	stop()                      //发停止指令
	wg.Wait()
}

func getUser(valCtx context.Context) {
	for {
		select {
		case <-valCtx.Done():
			fmt.Println("[获取用户]", "协程退出")
			return
		default:
			useId := valCtx.Value("UserId")
			fmt.Println("[获取用户]", "用户ID：", useId)
			time.Sleep(time.Second)
		}

	}
}

func watchDog(ctx context.Context, name string) {
	//开启 for select循环，一直后台监控
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "停止指令已收到，马上停止")
			return
		default:
			fmt.Println(name, "正在监控....")
		}
		time.Sleep(time.Second)
	}
}
