package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
)

func compareSHa256Byte(b1 [32]byte, b2 [32]byte) (num int) {
	counter := 0
	for i := 0; i < len(b1); i++ {
		if b1[i] != b2[i] {
			counter++
		}
	}
	return counter
}

func main() {
	s := "zhang san "
	c1 := sha256.Sum256([]byte(s))
	c2 := sha256.Sum256([]byte(s))
	fmt.Println(reflect.TypeOf([]byte(s)))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Println(len(c1), len(c2))
	fmt.Println(compareSHa256Byte(c1, c2))
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}
