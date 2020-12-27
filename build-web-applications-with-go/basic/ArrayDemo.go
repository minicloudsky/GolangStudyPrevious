package main

import "fmt"

func PrintfArray(arr [10]int) {
	for i, value := range arr {
		if i != len(arr)-1 {
			fmt.Printf(" %d ,", value)
		} else{
			fmt.Printf(" %d \n", value)
		}
	}
}

func myBubbleSort(arr *[10]int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			var temp int
			if arr[i] > arr[j] {
				temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
}
//func main() {
//	var arr = [10]int{45, 12, 6, 1, 21, 90, -21, 34, 45, 69}
//
//	fmt.Println("before sort: ")
//	PrintfArray(arr)
//	myBubbleSort(&arr)
//	fmt.Println("after sort: ")
//	PrintfArray(arr)
//}
