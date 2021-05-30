package main

import "fmt"

func main() {
	var arr [100]int
	s := arr[0:10]
	s = s[10:20]
	fmt.Println(cap(s))
	fmt.Println(len(s))
}
