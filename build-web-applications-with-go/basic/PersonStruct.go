package main

// 声明一个新的类型
type person struct {
	name string
	age  int
}
type student struct {
	name  string
	age   int
	score float32
}

// 比较两个人的年龄，返回年龄大的那个人，并且返回年龄差
// struct也是传值的
func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age { // 比较p1和p2这两个人的年龄
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func Higher(s1, s2 student) (student, float32) {
	if s1.score > s2.score {
		return s1, s1.score - s2.score
	}
	return s2, s2.score - s1.score
}

//func main() {
//	//var tom person
//	//
//	//// 赋值初始化
//	//tom.name, tom.age = "Tom", 18
//	//
//	//// 两个字段都写清楚的初始化
//	//bob := person{age: 25, name: "Bob"}
//	//
//	//// 按照struct定义顺序初始化值
//	//paul := person{"Paul", 43}
//	//
//	//tbOlder, tbDiff := Older(tom, bob)
//	//tpOlder, tpDiff := Older(tom, paul)
//	//bpOlder, bpDiff := Older(bob, paul)
//	//
//	//fmt.Printf("Of %s and %s, %s is older by %d years\n",
//	//	tom.name, bob.name, tbOlder.name, tbDiff)
//	//
//	//fmt.Printf("Of %s and %s, %s is older by %d years\n",
//	//	tom.name, paul.name, tpOlder.name, tpDiff)
//	//
//	//fmt.Printf("Of %s and %s, %s is older by %d years\n",
//	//	bob.name, paul.name, bpOlder.name, bpDiff)
//	var lily student
//	lily.name = "lily"
//	lily.age = 12
//	lily.score = 76.12
//	tom := student{score: 78, name: "tom", age: 12}
//	jack := student{"jack", 12, 89.12}
//	var higherUser, score = Higher(tom, jack)
//	fmt.Println(higherUser, score)
//}
