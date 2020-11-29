package main

import "fmt"

//var name string
//var age int64
//var salary float32
//var vname1, vname2, vname3 = 1, 2, 3

func main() {
	// 简短声明，只能在 func 内部
	//age1, age2, age3 := 1, 2, 3
	// _ 是特殊变量，赋值给它的变量都会被丢弃
	//_, b := 34, 35
	// 已声明但未使用的变量会在编译阶段报错
	// 常亮
	const name string = "minicloudsky"
	const age int32 = 23
	const prefix string = "mini*"
	//  bool 类型
	//var isActive bool
	//var enabled, disabled = true, false  // 忽略类型的声明
	//var available bool  // 一般声明
	//valid := false      // 简短声明
	//available = true    // 赋值操作
	// 数值类型
	// rune, int8, int16, int32, int64和byte, uint8, uint16, uint32, uint64。
	// 其中rune是int32的别称，byte是uint8的别称
	// 类型的变量之间不允许互相赋值或操作，不然会在编译时引起编译器报错
	// Go还支持复数。它的默认类型是complex128（64位实数+64位虚数）。
	//如果需要小一些的，也有complex64(32位实数+32位虚数)。
	//复数的形式为RE + IMi，其中RE是实数部分，IM是虚数部分，而最后的i是虚数单位
	//var c complex64 = 5 + 5i
	//fmt.Println(c)
	// 字符串，Go中的字符串都是采用UTF-8字符集编码。
	// 字符串是用一对双引号（""）或反引号（`  ` ）括起来定义，它的类型是string
	//shanghai, shenzhen, hangzhou := "shanghai", "shenzhen", "hangzhou"
	//fmt.Println(shanghai, shenzhen, hangzhou)
	// 声明一个字符串，初始化为空字符串
	//var emptyString string
	//emptyString = "yanjiyou"
	//fmt.Println(emptyString)
	// 字符串值不可变，需要改变可以这样实现
	//s := "hello"
	//c := []byte(s) // 将字符串 s 转换为 []byte 类型
	//fmt.Println("Before transfer: ")
	//fmt.Println(s)
	//fmt.Println(c)
	//c[0] = 'c'
	//s2 := string(c) // 再转换回 string 类型
	//fmt.Printf("%s\n", s2)
	//s := "hello"
	//s = "c" + s[1:] // 字符串虽不能更改，但可进行切片操作
	//fmt.Printf("%s\n", s)
	// + 可以做字符串连接
	//vivo, oppo := "vivo", "oppo"
	//var phone string
	//phone = vivo + " , " + oppo
	//fmt.Println("%s\n", phone)
	// 声明多行字符串
	// 括起的字符串为 Raw字符串，即字符串在代码中的形式就是打印时的形式，它没有字符转义，换行也将原样输出
	//m := `hello
	//world`
	//fmt.Println(m)
	// 错误类型
	//err := errors.New("emit macho dwarf: elf header corrupted")
	//if err != nil {
	//	fmt.Print(err)
	//}
	// 分组声明，同时声明多个常量、变量，或者导入多个包时，可采用分组的方式进行声明
	//import "fmt"
	//import "os"
	//
	//const i = 100
	//const pi = 3.1415
	//const prefix = "Go_"
	//
	//var i int
	//var pi float32
	//var prefix string
	//import(
	//	"fmt"
	//	"os"
	//)
	//
	//const(
	//	i = 100
	//	pi = 3.1415
	//	prefix = "Go_"
	//)
	//
	//var(
	//	i int
	//	pi float32
	//	prefix string
	//)
	// iota 声明enum的时候采用，它默认开始值是0，const中每增加一行加1
	//const (
	//	x = iota
	//	y = iota
	//	z = iota
	//	w // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
	//)
	//const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0
	//const (
	//	h, i, j = iota, iota, iota //h=0,i=0,j=0 iota在同一行值相同
	//)
	//const (
	//	a       = iota //a=0
	//	b       = "B"
	//	c       = iota             //c=2
	//	d, e, f = iota, iota, iota //d=3,e=3,f=3
	//	g       = iota             //g = 4
	//)
	//fmt.Println(a, b, c, d, e, f, g, h, i, j, x, y, z, w, v)
	/**
	除非被显式设置为其它值或iota，每个const分组的第一个常量被默认设置为它的0值，
	第二及后续的常量被默认设置为它前面那个常量的值，如果前面那个常量的值是iota，
	则它也被设置为iota
	*/
	// Go 程序设计规则
	// 大写字母开头的变量是可导出的，也就是其他包可以读取的，是公有变量；
	// 小写字母开头为不可导出，是私有变量
	// 大写字母开头的函数也是一样，相当于class中的带public关键词的公有函数；
	// 小写字母开头的就是有private关键词的私有函数。
	// array 数组
	// var arr[n]type
	//var arr [10]int
	//arr[0] = 42 // 下标从 0 开始
	//arr[1] = 45 // 赋值
	//fmt.Println("The first element is %d\n", arr[0])
	//fmt.Println("The last element is %d\n", arr[9])
	//长度也是数组类型的一部分，因此[3]int与[4]int是不同的类型，数组也就不能改变长度。
	//数组之间的赋值是值的赋值，即当把一个数组作为参数传入函数的时候，传入的其实是该数组的副本，而不是它的指针。
	//如果要使用指针，那么就需要用到后面介绍的slice类型
	//// 数组可以用 := 声明
	//id_arr := [3]int{1, 2, 3} // 声明长度 为 3 的 int 数组
	//fmt.Println(id_arr)
	//age_arr := [10]int{1, 2, 3} // 声明了一个长度为10的int数组，其中前三个元素初始化为1、2、3，其它默认为0
	//fmt.Println(age_arr)
	//token_arr := [...]int{4, 5, 6} // 省略长度用 ... 代替，根据元素个数计算数组长度
	//fmt.Println(len(token_arr))
	// 多维数组
	// 声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中又有4个int类型的元素
	//doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	//fmt.Println(doubleArray)
	//// 上面的声明可以简化，直接忽略内部的类型
	//easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	//fmt.Println(easyArray)
	// slice 动态数组
	// slice并不是真正意义上的动态数组，而是一个引用类型。
	//slice总是指向一个底层array，slice的声明也可以像array一样，只是不需要长度
	//var fslice []int
	//fmt.Println(reflect.TypeOf(fslice))
	//slice := []byte{'a', 'b', 'c', 'd'}
	//fmt.Println(slice)
	//var ar = [10]byte{'a', 'b'}
	//fmt.Println(ar)
	//// 声明两个含有byte的slice
	//var a, b []byte
	//
	//// a指向数组的第3个元素开始，并到第五个元素结束，
	//a = ar[2:5]
	////现在a含有的元素: ar[2]、ar[3]和ar[4]
	//fmt.Println(a)
	//// b是数组ar的另一个slice
	//b = ar[3:5]
	//// b的元素是：ar[3]和ar[4]
	//fmt.Println(b)
	//// slice的默认开始位置是0，ar[:n]等价于ar[0:n]
	//// slice的第二个序列默认是数组的长度，ar[n:]等价于ar[n:len(ar)]
	//// 如果从一个数组里面直接获取slice，可以这样ar[:]，因为默认第一个序列是0，第二个是数组的长度，即等价于ar[0:len(ar)]
	//// 声明数组
	//var future = [10]byte{'b', 'y', 't', 'e', 'd', 'a', 'n', 'c', 'e'}
	//// 声明两个 slice
	//var aSlice, bSlice []byte
	//aSlice = future[:3]
	//aSlice = future[5:]
	//aSlice = future[:]
	//// 从 slice 中获取 slice
	//aSlice = future[3:7]
	//bSlice = aSlice[1:3]
	//bSlice = aSlice[:3]
	//bSlice = aSlice[0:5]
	//bSlice = aSlice[:]
	//fmt.Println(aSlice)
	//fmt.Println(bSlice)
	//slice是引用类型，所以当引用改变其中元素的值时，其它的所有引用都会改变该值，例如上面的aSlice和bSlice，
	//如果修改了aSlice中元素的值，那么bSlice相对应的值也会改变。
	/**
	slice 更像一个结构体，包含三个元素
	一个指针，指向数组中 slice 指定的开始位置
	长度，即 slice 的长度
	最大长度，也就是 slice 开始位置到数组的最后位置的长度
	*/
	//fmt.Println(len(aSlice))         // len 获取 slice 长度
	//fmt.Println(cap(aSlice))         // cap 获取 slice 的最大容量
	//newaSlice := append(aSlice, 255) // 追加一个或多个元素，返回和slice一样类型的slice
	//fmt.Println("newSlice : ", newaSlice)
	//var copySlice []byte
	//copySlice = bSlice[:]
	//fmt.Println("bSlice : ", bSlice)
	//fmt.Println("before copySlice: ", copySlice)
	//copyNum := copy(newaSlice, copySlice) // copy 从 源 slice 的 src 中复制元素到目标dst，并且返回复制的元素的个数
	//
	//fmt.Println("copySlice: ", copySlice)
	//fmt.Println("copy num:", copyNum)
	/*var array [10]int
	// 容量是 8
	slice := array[2:4]
	fmt.Println("array capacity: ", cap(array))
	fmt.Println("slice capacity: ", cap(slice))*/
	// 指定容量，7-2 = 5 这样这个产生的新的slice就没办法访问最后的三个元素
	// 如果slice是这样的形式array[:i:j]，即第一个参数为空，默认值就是0。
	//slice = array[2:4:7]
	// map 类似 python 中的 dict map[keyType]valueType
	// 声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
	//var numbers map[string]int
	//fmt.Println(numbers)
	// 另一种 map 的声明方式
	users := make(map[string]int)
	users["tom"] = 1
	users["jerry"] = 10
	users["root"] = 100
	fmt.Println(users)
	fmt.Println(users["tom"])
	fmt.Println("users length: ", len(users))
	delete(users, "tom")
	fmt.Println("after delete key: ", users)
	/*
		map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取
		map的长度是不固定的，也就是和slice一样，也是一种引用类型
		内置的len函数同样适用于map，返回map拥有的key的数量
		map的值可以很方便的修改，通过numbers["one"]=11可以很容易的把key为one的字典值改为11
		map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制
	*/
	//rating := map[string]float32{"java": 3.2, "python": 12.3, "nodejs": 34.5}
	//csharpRating, ok := rating["c#"]
	//fmt.Println(csharpRating, ok)
	//if ok {
	//	fmt.Println("C# is in the rating and it's value is: ", csharpRating)
	//} else {
	//	fmt.Println("We have no rating about csharp")
	//}
	//delete(rating, "java")
	//fmt.Println("after delete rating: ", rating)
	m := make(map[string]string)
	m["Hello"] = "World"
	fmt.Println("before m:", m)
	m1 := m
	m1["Hello"] = "Salut" // 现在m["hello"]的值已经是Salut了
	fmt.Println("after m: ", m)
	fmt.Println("m1: ", m1)
}
