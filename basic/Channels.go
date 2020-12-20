package main

import "fmt"

func main() {
	//var name string
	//var age int
	//name = "tony"
	//age = 22
	ci := make(chan int)
	cs := make(chan string)
	cf := make(chan interface{})
	fmt.Println(ci, cs, cf)
}
