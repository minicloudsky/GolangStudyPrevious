package main

import "fmt"

func main(){
    slice1 := make([]float32,0)
    slice2 := make([]float32,0,5)
    fmt.Println(slice1)
    fmt.Println(len(slice2),cap(slice2))
    slice2 = append(slice2,1,2,3,4,5)
    fmt.Println(len(slice2),cap(slice2))
    fmt.Println("slice2: ",slice2)
    // 子切片 [start,end]
/*
声明切片时可以为切片设置容量大小，为切片预分配空间。
在实际使用的过程中，如果容量不够，切片容量会自动扩展。
*/
    sub1 := slice2[3:] // [4,5]
    fmt.Println("slice2[3:] ",sub1)
    sub2 := slice2[:3] // [0,1,2]
    sub3 := slice2[1:4]  // [2,3,4]
    fmt.Println("slice[:3] ", sub2)
    fmt.Println("slice[1:4] ", sub3)
}

