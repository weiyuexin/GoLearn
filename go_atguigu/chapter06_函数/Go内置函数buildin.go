package main

import "fmt"

//todo 内置函数
func main() {
	// golang提供的一些常用的函数，不需要导包就可以使用
	// 1、len() 求变量或常量所占内存长度，字节为单位
	// 2、new() 用于分配内存，主要分配值类型
	// 3、make() 用于分配内存，主要分配引用类型,比如map、slice、channel等

	// new()的使用
	num := 100
	fmt.Printf("num的类型%T , num的值=%v ，num的地址%v\n", num, num, &num) //num的类型int , num的值=100 ，num的地址0xc0000aa058

	num1 := new(int) //*int
	// num1的类型：*int
	// num1的值：地址 0xc0000aa090
	// num1的地址：地址 0xc0000ce020
	fmt.Printf("num1的类型%T , num1的值=%v ，num1的地址%v num1这个指针指向的值=%v\n", num1, num1, &num1, *num1) //num1的类型*int , num1的值=0xc0000aa090 ，num1的地址0xc0000ce020

}
