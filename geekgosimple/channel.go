package main

import "fmt"
import "time"

var ch =  make(chan string, 10)

func download(url string){
    fmt.Println("start to download ",url)
    time.Sleep(time.Second)
    ch <- url // 将 url发送给信道
}

func main(){
    for i:=0;i<3;i++{
        go download("a.com/" + string(i+'0'))
    }
    for i:=0; i <3;i++{
        msg := <-ch // 等待信道返回消息
        fmt.Println("finish", msg)
    }
    fmt.Println("Done!")
}

