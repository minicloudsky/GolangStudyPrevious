package main

import "fmt"

func main(){
    nums := []int{10,20,30,40}
    for i,num := range nums{
        fmt.Println(i,num)
    }
    m2 := map[string]string{
        "Sam":"Male",
        "Alice":"Female",
    }
    for key,value := range m2{
        fmt.Println(key, value)
    }
}

