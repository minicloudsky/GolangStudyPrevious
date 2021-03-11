package main

import "fmt"

// type 关键字可以定义新类型
type Gender int8
const (
    MALE Gender = 1
    FEMALE Gender = 2
)
var gender = MALE

func main(){
    switch gender {
    case FEMALE:
        fmt.Println("female")
    case MALE:
        fmt.Println("male")
    default:
        fmt.Println("unknown")
    }
}

