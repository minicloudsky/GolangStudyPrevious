package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "jia ya wu"
	name2 := strings.Split(name, " ")
	for i := 0; i < len(name2); i += 1 {
		fmt.Println(name2[i])
	}
	var arr [10]string
	var aww [10]int
	var slice1 = make([]int, 100)
	var slice4 = make([]string, 100)
	var countryCapitalMap map[string]string
	m1 := make(map[string]string)

	fmt.Println(arr)
	fmt.Println(slice1)

	fmt.Println("go")
	ch := make(chan int)
	fmt.Println(ch)
}

type user interface {
}
