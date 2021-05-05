package main

import "fmt"
import "reflect"

func main(){
    var arr [10]int
    var slice2  []int
    fmt.Println(reflect.TypeOf(arr))
    fmt.Println(reflect.TypeOf(slice2))
}

