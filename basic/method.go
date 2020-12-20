package main
//
//import (
//	"fmt"
//	"math"
//)
//
//type Rectangle struct {
//	width, height float64
//}
//
//type Circle struct {
//	radius float64
//}
//
//func (r Rectangle) area() float64 {
//	return r.width * r.height
//}
//func area(r Rectangle) float64 {
//	return r.height * r.width
//}
//
//// method是附属在一个给定的类型上的，他的语法和函数的声明语法几乎一样，
////只是在func后面增加了一个receiver(也就是method所依从的主体)
//// 计算 Rectangle 结构体的周长，name 为参数，，接受一个 Rectangle 结构体为 Receiver
//// 虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样
//// method里面可以访问接收者的字段
//// 调用 method 通过.访问，就像 struct 里面访问字段一样
//func (r Rectangle) perimeter(name string) float64 {
//	fmt.Println("this is ", name)
//	return (r.height + r.width) * 2
//}
//func (c Circle) perimeter() float64 {
//	return 2 * math.Pi * c.radius
//}
//
//// 普通的 function，参数为Circle 结构体，返回一个 float64 类型的面积
//func (c Circle) area() float64 {
//	return c.radius * c.radius * math.Pi
//}
//
//func main() {
//	r1 := Rectangle{12, 2}
//	r2 := Rectangle{9, 4}
//	c1 := Circle{10}
//	c2 := Circle{25}
//
//	fmt.Println("Area of r1 is: ", r1.area())
//	fmt.Println("Area of r2 is: ", r2.area())
//	fmt.Println("Area of c1 is: ", c1.area())
//	fmt.Println("Area of c2 is: ", c2.area())
//	fmt.Println("perimeter of r1 is :", r1.perimeter("r1"))
//	fmt.Println("perimeter of r2 is: ", r2.perimeter("r2"))
//	fmt.Println("perimeter of c1 is: ", c1.perimeter())
//	fmt.Println("perimeter of c2 is : ", c2.perimeter())
//	// 自定义类型
//	type user string
//	var employee user
//	employee = "zhangsan"
//	fmt.Println(employee)
//}
