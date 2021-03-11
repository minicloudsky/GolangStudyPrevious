package main

import "fmt"

/*

Go 也有 private public的概念，粒度是包，如果类型/接口/方法/函数/字段的首字母大写，则是
public的，对其他包可见，如果首字母小姐，则是Private的，对其他包不可见
*/
func main(){
    fmt.Println(add(3,5))
}

