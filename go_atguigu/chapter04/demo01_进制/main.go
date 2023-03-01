package main

import "fmt"

// TODO 进制
func main() {
	// 对应整数，有以下四种方式
	// 1、二进制
	//		在golang中，不能直接使用二进制表示一个整数，他沿用了c的特点
	//		%b 表示按二进制输出
	var i int = 5
	fmt.Printf("二进制:%b\n", i) //二进制:101
	// 2、十进制
	// 3、八进制
	//		在golang中，八进制可以直接表示, 以0开头
	var j int = 011
	fmt.Println("j=", j) //j= 9
	// 4、十六进制
	// 		在golang中，十六进制可以直接表示，以0x开头
	var k int = 0x11
	fmt.Println("k=", k) //k= 17
}
