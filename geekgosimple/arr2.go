package main

import "fmt"

func main(){
    // 声明数组
    var arr [5]int
    var arr2 [5][5]int
    arr[0] = 12
    arr2[0][0] = 123
    fmt.Println(arr)
    fmt.Println(arr2)
    // 声明时候初始化
    var arr3 = [5]int{1,2,3,4,5}
    arr4 := [3]int{11,22,33}
    arr = [5]int{1,2,3,4,5}
    for i:=0;i < len(arr);i++{
        arr[i] += 100
    }
    fmt.Println(arr3)
    fmt.Println(arr4)
    

}

