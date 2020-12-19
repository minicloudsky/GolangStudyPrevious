package main

type Person struct {
	name string
	age  int
}
//
//func main() {
//	var p Person
//	// 按照顺序提供初始化值
//	p.name = "zhangsan"
//	p.age = 123
//	fmt.Println(p, p.name, p.age)
//	// 按照顺序提供初始化值
//	p2 := Person{"tony", 32}
//	fmt.Println(p2)
//	// 通过field:value的方式初始化，这样可以任意顺序
//	p3 := Person{age: 12, name: "stone"}
//	fmt.Println(p3)
//	pointer := new(Person)
//	pointer.name = "jack"
//	pointer.age = 12
//	fmt.Println(pointer)
//
//}
