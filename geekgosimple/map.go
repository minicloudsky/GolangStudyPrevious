package main

import "fmt"

func main(){
    // 声明
    m1 := make(map[string]int)
    m2 := map[string]string{
        "Sam": "Male",
        "Alice": "Female",
    }
    // 赋值 // 修改
    m1["Tom"] = 18
    fmt.Println("m1: ",m1)
    fmt.Println("m2: ",m2)
}

