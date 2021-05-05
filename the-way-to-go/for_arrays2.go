package main

import "fmt"

func main(){
    var arr [10]int
    for i:=0; i< len(arr);i++{
        arr[i] = i*3
    }
    for i,val := range arr {
        fmt.Printf("index %d val %d\n", i, val)
    }
}

