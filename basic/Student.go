package main
//
//import (
//	"fmt"
//	"log"
//	"math"
//)
//
//const (
//	SHANGHAI = iota
//	ZHEJIANG
//	JIANGSU
//	JIANGXI
//	HUNAN
//)
//
//type Person struct {
//	name, sex string
//	age       int
//}
//type Address struct {
//	province, city string
//}
//type Student struct {
//	Person
//	Address
//	school, class string
//	balance       float64
//}
//type studentList []Student
//
//func (s Student) isRich(name string) bool {
//	if s.province == "guangdong" && s.city == "shenzhen" && s.balance > 1000 {
//		log.Println(name, "is rich .")
//		return true
//	}
//	log.Println(name, "is poor .")
//	return false
//}
//func (s *Student) enRich(name string) {
//	s.balance = s.balance * math.Pi
//	log.Println(s, name, "becones rich ~")
//}
//func (s studentList) LiveInShanghai() studentList {
//	for index, _ := range s {
//		s[index].province = "shanghai"
//	}
//	return s
//}
//
//func main() {
//	var tom = Student{
//		Person:  Person{name: "tom", sex: "male", age: 23},
//		Address: Address{province: "guangdong", city: "shenzhen"},
//		school:  "Shenzhen University",
//		class:   "biology",
//		balance: 123.92,
//	}
//	fmt.Println(tom)
//	fmt.Println(tom.isRich(tom.name))
//	(&tom).enRich(tom.name)
//	fmt.Println("after enrich,tom: ", tom)
//	//fmt.Println("after enrich,tom:", tom)
//	//fmt.Println(tom.balance)
//	//var students studentList
//	//students = studentList{Student{
//	//	Person:  Person{name: "jack", age: 12},
//	//	Address: Address{province: "sichuan", city: "chengdu"},
//	//	school:  "chengdu universiy",
//	//	class:   "math",
//	//	balance: 3435.34,
//	//}, Student{
//	//	Person:  Person{name: "lucy", age: 23},
//	//	Address: Address{province: "sichuan", city: "chengdu"},
//	//	school:  "chengdu universiy",
//	//	class:   "math",
//	//	balance: 3435.34,
//	//}, Student{
//	//	Person:  Person{name: "tony", age: 45},
//	//	Address: Address{province: "sichuan", city: "chengdu"},
//	//	school:  "chengdu universiy",
//	//	class:   "math",
//	//	balance: 3435.34,
//	//}}
//	//for index, value := range students {
//	//	fmt.Println(index, value)
//	//}
//	//fmt.Println(HUNAN)
//	//students = students.LiveInShanghai()
//	//fmt.Println(students)
//}
