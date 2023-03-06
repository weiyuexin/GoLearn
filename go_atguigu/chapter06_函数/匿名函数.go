package main

import "fmt"

//定义全局匿名函数
var (
	fun1 = func(n1, n2 int) int {
		return n1 + n2
	}
)

//todo 匿名函数
func main() {
	//匿名函数的使用方法
	//1、定义时直接调用（只能用一次）
	//2、将匿名函数赋值给一个变量，通过变量调用函数（可以多次调用）
	//3、将匿名函数赋值给一个全局变量，就可以实现全局调用

	//使用匿名函数求两个数的和
	res1 := func(n1, n2 int) int {
		return n1 + n2
	}(2, 3)
	fmt.Println("res1=", res1)

	//将匿名函数赋值给f,此时可以使用f调用函数
	f := func(n1, n2 int) int {
		return n1 + n2
	}
	res2 := f(20, 30)
	fmt.Println("res2=", res2)

	//调用全局匿名函数
	res3 := fun1(3, 4)
	fmt.Println("res3=", res3)
}
