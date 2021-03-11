package main

import "fmt"

func test(){
    fmt.Println(1)
    defer func(){
        fmt.Println(2)
    }
    fmt.Println(3)
}

func main(){
    test()
}
