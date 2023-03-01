package main

import "fmt"

func main() {
	// TODO 演示 & 和 *
	a := 10
	fmt.Println("a的地址:", &a) //0xc0000aa058
	var ptr *int = &a
	fmt.Println("ptr 指向的值是:", *ptr) //ptr 指向的值是: 10

	// go中没有三元运算符，可以使用if-else实现
	var n int
	var i, j int = 10, 12
	// n = i > j ? i : j
	if i > j {
		n = i
	} else {
		n = j
	}
	fmt.Println("n=", n) //n= 12

}
