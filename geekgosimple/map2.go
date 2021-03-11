package main

import "fmt"

func main(){
    var phone = make(map[string]int)
    phone["tom"] = 110
    phone["jack"] = 123
    book := make(map[string]int)
    book["php"] = 34
    book["java"] = 45
    book["python"] = 90
    computer := map[string]string{
        "dell":"USA",
        "ACER":"Taiwan",
        "Lenovo":"China",
        "Samsung":"Korea",
    }
    fmt.Println("book: ", book)
    fmt.Println("phone: ", phone)
    fmt.Println("computer: ", computer)
}

