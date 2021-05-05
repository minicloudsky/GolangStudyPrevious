package main

import "fmt"
import "reflect"

func main(){
    var arrKeyValue = [...]string{3:"Chris",4:"Ron"}
    for i:=0; i< len(arrKeyValue);i++{
        fmt.Printf("Person at %d is %s\n",i,arrKeyValue[i])
    }
    fmt.Println(reflect.TypeOf(arrKeyValue))
}

