package main

import "fmt"
import "os"

func main() {
	var (
		HOME   = os.Getenv("HOME")
		USER   = os.Getenv("USER")
		GOROOT = os.Getenv("GOROOT")
	)
	fmt.Println(HOME, USER, GOROOT)
}
