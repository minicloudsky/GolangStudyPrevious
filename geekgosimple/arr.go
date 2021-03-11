package main

import "fmt"

func main(){
    var arr = [5]int {1,2,3,4,5} // 一维数组
    fmt.Println(arr)
    var arr2 [5][5] int // 二维
    for i:=0; i<len(arr2); i++{
        for j :=0; j<len(arr2);j++{
            arr2[i][j] = i+j
        }
    }
    for i:=0; i<len(arr2); i++{
        for j :=0; j< len(arr2);j++{
            fmt.Print(arr2[i][j]," ")
        }
        fmt.Println("")
    }
}

