package main

import "fmt"
import "github.com/jinzhu/copier"

func main(){
    arr :=  [...]int{1,3,5,6,7,9} 
    slice := arr[2:4]
    var deepCopyValue []int
    copier.Copy(&slice, deepCopyValue)
    fmt.Println("arr: ", arr)
    fmt.Println("slice: ", slice)
    fmt.Println("deepCopyValue: ", deepCopyValue)
}
