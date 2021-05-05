package main

import "fmt"

func main(){
    fmt.Println("Hello, World!")
    var arrKeyValue = [5]string{3:"zhangsan",4:"lisi"}
    for i:=0;i<len(arrKeyValue); i++ {
        fmt.Printf("Person at %d is %d\n",i,arrKeyValue[i])
    }
}

