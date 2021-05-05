package main

import "fmt"
import "reflect"

func main(){
    arr := [10]int{1,2,3,4,5,6,7,8,9,10}
    var slice1 = arr[2:7]
    s := []int{1,2,3}
    fmt.Println(s)
    fmt.Println(slice1)
    fmt.Println("type arr ",reflect.TypeOf(arr))
    fmt.Println("type slice1 ",reflect.TypeOf(slice1))
    fmt.Printf("capacity of slice1: %d\n",cap(slice1))
    fmt.Printf("len of slice1: %d\n",len(slice1))
}

