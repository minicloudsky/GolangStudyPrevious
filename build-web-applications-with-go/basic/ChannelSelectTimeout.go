package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				println(v)
			case <-time.After(5 * time.Second):
				println("timeout")
				o <- true
				break
			}
		}
	}()
	<-o
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
}

//Goexit
//
//退出当前执行的goroutine，但是defer函数还会继续调用
//
//Gosched
//
//让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行。
//
//NumCPU
//
//返回 CPU 核数量
//
//NumGoroutine
//
//返回正在执行和排队的任务总数
//
//GOMAXPROCS
//
//用来设置可以并行计算的CPU核数的最大值，并返回之前的值。
