package main

import (
	"fmt"
	"math/big"
)

var name string

func main() {

	var num = new(big.Int)
	num.SetString("240000000000000000000000000", 10)
	fmt.Printf("%T", num)

	//var count int = 0
	//for count = 0; count < 10; count++ {
	//	fmt.Println("for 循环之内: ", count)
	//}
	//
	//fmt.Println("for 循环外： ", count)
	var n int = 10
	test(n)
	fmt.Println("当前 n 的值有变化了吗？ ", n)

	//var s = '你'
	//fmt.Printf("%#v,   %T", s, s)

	var s1 = "s"
	fmt.Printf("%#v,   %T", s1, s1)

	var s2 = 's'
	fmt.Printf("  %#v,   %T", s2, s2)

	var blank string
	fmt.Printf("   %#v,   %T", blank, blank)
}

func test(n int) {
	n += 1
	fmt.Println("当前 n 的值是： ", n)
}
