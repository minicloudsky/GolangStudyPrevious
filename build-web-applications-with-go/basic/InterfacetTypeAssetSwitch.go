package main
//
//import (
//	"fmt"
//	"reflect"
//	"strconv"
//)
//
//type Element interface {
//}
//
//// go  interface 支持类似于 struct 的嵌入字段
//type NewElement interface {
//	Element
//}
//type List []Element
//
//type Person struct {
//	name string
//	age  int
//}
//
////打印
//func (p Person) String() string {
//	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
//}
//
//func main() {
//	list := make(List, 3)
//	list[0] = 1       //an int
//	list[1] = "Hello" //a string
//	list[2] = Person{"Dennis", 70}
//	//list_type, ok := list[0].(int)
//	//fmt.Println(list_type, ok)
//	//for index, element := range list {
//	//	switch value := element.(type) {
//	//	case int:
//	//		fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
//	//	case string:
//	//		fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
//	//	case Person:
//	//		fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
//	//	default:
//	//		fmt.Println("list[%d] is of a different type", index)
//	//	}
//	//}
//	// 反射
//	//var name string
//	//name = "hello"
//	//t := reflect.TypeOf(list[2]) //得到类型的元数据,通过t我们能获取类型定义里面的所有元素
//	//v := reflect.ValueOf(t)      //得到实际的值，通过v我们获取存储在里面的值，还可以去改变值
//	//fmt.Println(t)
//	//fmt.Println(v)
//
//	//tag := t.Elem().Field(0).Tag      //获取定义在struct里面的标签
//	//name = v.Elem().Field(0).String() //获取存储在第一个字段里面的值
//	//fmt.Println(tag)
//	//fmt.Println(name)
//	//var x float64 = 3.4
//	//v := reflect.ValueOf(x)
//	//fmt.Println("type:", v.Type())
//	//fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
//	//fmt.Println("value:", v.Float())
//	var x float64 = 3.4
//	p := reflect.ValueOf(&x)
//	v := p.Elem()
//	v.SetFloat(7.1)
//	fmt.Println(x)
//	fmt.Println(p)
//	fmt.Println(v)
//}
