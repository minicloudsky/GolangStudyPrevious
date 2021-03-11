package main

import "fmt"

func main(){
    m := make(map[string]interface{})
    m["name"] = "Tom"
    m["age"] = 18
    m["scores"] = [3]int{98,99,100}
    fmt.Println(m) // 
}

