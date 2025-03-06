package main

import (
	"fmt"
	"strings"
)

var name string

type crater struct {
}

func (c crater) String() string {
	return "我是一个陨石"
}

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type newInt int

func reclassify(planets []string) {
	planets = planets[0:8]
	planets[0] = "new1"
	fmt.Println("函数中的打印: ", planets)
	fmt.Printf("plants在函数中获取到的地址:%p\n", &planets)
}

func main() {

	crater := crater{}
	fmt.Println(crater)

	//var p *int
	//v := p
	//fmt.Printf("看看类型： %T %v", v, v)

	//for value := range 10 {
	//	fmt.Println("啦啦啦啦...", value)
	//}

	//planets := []string{
	//	"Mercury", "Venus", "Earth", "Mars",
	//	"Jupiter", "Saturn", "Uranus", "Neptune",
	//	"Pluto",
	//}
	//reclassify(planets)
	//fmt.Printf("plants在函数外获取到的地址:%p\n", &planets)
	//fmt.Println("main中打印： ", planets)

	//l := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//s1 := l[1:5]
	//s2 := l[:5]
	//fmt.Printf("s1在函数外获取到的地址:%p\n", &s1)
	//fmt.Printf("s2在函数外获取到的地址:%p\n", &s2)
	//
	//fmt.Println(s1, s2)
	//fmt.Println(s1[0], s2[0])
	//s1[0] = 100
	//fmt.Println(s1, s2)

	//na := newInt(10)
	//fmt.Println(na)
	//shout(crater{})

	//var num = new(big.Int)
	//num.SetString("240000000000000000000000000", 10)
	//fmt.Printf("%T", num)
	//
	////var count int = 0
	////for count = 0; count < 10; count++ {
	////	fmt.Println("for 循环之内: ", count)
	////}
	////
	////fmt.Println("for 循环外： ", count)
	//var n int = 10
	//test(n)
	//fmt.Println("当前 n 的值有变化了吗？ ", n)
	//
	////var s = '你'
	////fmt.Printf("%#v,   %T", s, s)
	//
	//var s1 = "s"
	//fmt.Printf("%#v,   %T", s1, s1)
	//
	//var s2 = 's'
	//fmt.Printf("  %#v,   %T", s2, s2)
	//
	//var blank string
	//fmt.Printf("   %#v,   %T", blank, blank)
}

func test(n int) {
	n += 1
	fmt.Println("当前 n 的值是： ", n)
}
