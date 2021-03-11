package main

import "fmt"

func add(num1 int,num2 int) int{
    return num1 + num2
}

// 多返回值
func div(num1 int,num2 int)(int,int){
    return num1 / num2, num1 % num2
}
// 给返回值命名，简化 return
func dec(num1 int,num2 int)(ans int){
    ans = num1 - num2
    return 
}

// main 函数为程序执行入口
func main(){
    fmt.Println("23+23=",add(23,23))
    quo,rem := div(100,17)
    fmt.Println(quo, rem)
    fmt.Println(add(11,22))
    var age1 int
    var age2 int
    age1 = 78
    age2 = 12
    fmt.Println(dec(age1,age2))
}


