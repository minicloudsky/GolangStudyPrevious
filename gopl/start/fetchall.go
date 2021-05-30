package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

//func main() {
//	start := time.Now()
//	ch := make(chan string)
//	urls := []string{"http://baidu.com",
//		"http://zhihu.com", "http://sina.com", "http://sougou.com"}
//	for _, url := range urls {
//		go fetch(url, ch)
//	}
//	for range urls {
//		fmt.Println(<-ch)
//	}
//	fmt.Printf("%2.fs elapsed\n", time.Since(start).Seconds())
//}
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	fmt.Println(resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%2.fs %7d %s", secs, nbytes, url)
}
