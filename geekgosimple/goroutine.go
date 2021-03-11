package main

import (
    "fmt"
    "sync"
    "time"
    "reflect"
)

var wg sync.WaitGroup

func download(url string){
    fmt.Println("start to download", url)
    time.Sleep(time.Second)
    wg.Done()
}

func main(){
    fmt.Println(time.Time())
    for i :=0 ; i <3;i++{
        wg.Add(1)
        go download("a.com/" + string(i + '0'))
    }
    wg.Wait()
    fmt.Println("Done!")
    fmt.Println(time.Second)
    fmt.Println(time.Time())
}

