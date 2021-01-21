// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
	// "strings"
)

//!+
func main() {
	// var s,sep string
	startTime := time.Time()
	for i := 0; i < len(os.Args); i++ {
		fmt.Println("index: ", i)
		fmt.Println("value: ", os.Args[i])
	}
	endTime := time.Time()
	fmt.Println(startTime)
	fmt.Println(endTime)
	// fmt.Println("time cost: ", endTime-startTime)

}

//!-
