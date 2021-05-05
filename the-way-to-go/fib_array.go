package main

import "fmt"


func fib(n int) (int){
    if n <= 1{
        return 1
    }
    if n==2 {
        return 1
    }
    return fib(n-1) + fib(n-2)

}

func main(){
    for i:=1;i<=40;i++ {
        fmt.Printf("fib %d is %d\n", i, fib(i))
    }
}

