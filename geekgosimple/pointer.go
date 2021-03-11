package main

import "fmt"

func main(){
    str := "Golang"
    var p *string = &str
    *p = "Hello"
    fmt.Println(str)
}

