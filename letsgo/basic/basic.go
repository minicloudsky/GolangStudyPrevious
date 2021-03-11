package main

import (
	"fmt"
	"math"
	"strconv"
)

func main(){
    var i int64
    var b string
    fmt.Println("i is ",i)
    fmt.Println("b is ",b)
    var floatNum float64 = 1.0
    var price1,price2 float64 = 8.8,9.6
    fmt.Println(floatNum,price1,price2)
    ii := 1
    s := "Hello Go!"
    fmt.Println("ii is ", ii)
    fmt.Println("s is  ",s)
    sint := strconv.Itoa(97)
    fmt.Println(sint,sint == "97")
    var b1 bool
    b2 := false
    b3 := true
    if b1{
        fmt.Println("b1 is true")
    }
    if b2{
        fmt.Println("b2 is true")
    }
    if b3{
        fmt.Println("b3 is true")
    }
    var i64 int64
    i64 = 10
    fmt.Println(i64 + 10)
    i32 := int32(42)
    fmt.Println(i32 + 10)
    fmt.Println(
        math.MaxInt64,
    )
    
}

