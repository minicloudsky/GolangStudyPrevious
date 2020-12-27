package main

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

//func main() {
//	a := []int{7, 2, 8, -9, 4, 0}
//
//	c := make(chan int)
//	go sum(a[:len(a)/2], c)
//	go sum(a[len(a)/2:], c)
//	fmt.Println(a[:len(a)/2])
//	x, y := <-c, <-c // receive from c
//
//	fmt.Println(x, y, x+y)
//	fmt.Println(c)
//}
