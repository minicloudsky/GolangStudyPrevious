package main

import "fmt"
import "reflect"

func main(){
    str2 := "Go语言"
    runeArr := []rune(str2)
    fmt.Println(reflect.TypeOf(runeArr[2]).Kind()) // int32
    fmt.Println(runeArr[2], string(runeArr[2]))    // 35821 语
    fmt.Println("len(runeArr)：", len(runeArr))    // len(runeArr)： 4
}
