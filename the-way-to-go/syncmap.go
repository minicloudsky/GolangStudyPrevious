package main

import "fmt"
import "sync"

func main() {
	var m sync.Map
	m.Store("qcrao", 18)
	m.Store("stefino", 20)
	age, _ := m.Load("qcrao")
	fmt.Println(age.(int))

	m.Range(func(key, value interface{}) bool {
		name := key.(string)
		age := value.(int)
		fmt.Println(name, age)
		return true
	})
	m.Delete("qcrao")
	age, _  = m.Load("qcrao")
	fmt.Println(age)
}
