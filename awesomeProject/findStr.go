package main

import "fmt"

func findStr(s1 string, s2 string) int {
	lenS1 := len(s1)
	lenS2 := len(s2)
	equalsLen := 0
	for i := 0; i < lenS1-1; {
		for j := 0; j < lenS2; {
			if equalsLen == lenS2 {
				return i
			}
			if s1[i] == s2[j] {
				i += 1
				j += 1
				equalsLen += 1
			} else {
				j += 1
			}
		}
	}
	return -1
}
func main() {
	s1 := "nameaaadjfhdjsfhzhangsanzzz"
	s2 := "zhangsan"
	fmt.Println(findStr(s1, s2))
}
