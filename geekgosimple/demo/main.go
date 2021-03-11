package main

import "fmt"
import "rsc.io/quote"
import "example/calc"

func main(){
    fmt.Println(quote.Hello())
    fmt.Println(calc.Add(10,3))
}

