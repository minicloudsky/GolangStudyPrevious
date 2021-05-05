package main

import "fmt"

func main(){
    var arr [15]int
    for i:=0; i<len(arr); i++{
        arr[i] = i
    }
    for i,value := range arr {
        fmt.Printf("index: %d value: %d\n", i, value)
    }
}

