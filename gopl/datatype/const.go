package main

import (
	"fmt"
	"time"
)

const (
	a = iota
	b
	c
	d
	e
)

func main() {
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)     // "time.Duration 0"
	fmt.Printf("%T %[1]v\n", timeout)     // "time.Duration 5m0s"
	fmt.Printf("%T %[1]v\n", time.Minute) // "time.Duration 1m0s"
	fmt.Println(a, b, c, d, e)
}
