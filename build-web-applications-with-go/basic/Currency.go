package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

var balance int

func addBalance(money int) {
	for i := 0; i < 10; i++ {
		balance = balance + money
		runtime.Gosched()
		fmt.Println("now,balance is ", balance)
	}
}
//func main() {
//	runtime.GOMAXPROCS(6)
//	go addBalance(10)
//	go say("world") // 开一个新的Goroutines执行
//	say("hello")    // 当前Goroutines执行
//}

// 以上程序执行后将输出：
// hello
// world
// hello
// world
// hello
// world
// hello
// world
// hello
