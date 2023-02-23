package main

import "fmt"
// 全局变量
var m = 100

func foo()(int, string){
	return 10,"weiyuexin"
}

func main()  {
	var name string
	var age int = 1
	var isOk bool = true
	name = "weiyuexin"
	fmt.Println("hello "+name)
	fmt.Println(age)
	fmt.Println(isOk)
	// 批量声明
	var(
		a string = "a"
		b int = 2
		c bool = false
		d float32 = 1
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	// 一次初始化多个变量
	var name1,age1 = "wyx",23
	fmt.Println(name1)
	fmt.Println(age1)
	// 短变量声明：在函数内部可以使用 := 方式声明并初始化变量
	n := 101
	fmt.Println(n)

	// 匿名变量：在多重赋值时，如果想要忽略某个值，可以使用 匿名变量 ，匿名变量使用下划线_表示
	x,_ := foo()
	_,y := foo()
	fmt.Println("x=",x)
	fmt.Println("y=",y)
	/*
	函数外的每个语句都必须以关键字开始（var、const、func等）
	:=不能使用在函数外。
	_多用于占位，表示忽略值。
	*/

	// const 声明常量
	const pi = 3.1415926
	const e = 2.7182
	// 一起声明时，如果省略了值，则表示和上一行相同
	const(
		n1 = 100
		n2
		n3
	)
	fmt.Println(pi)
	fmt.Println(e)
	fmt.Println(n3) //100

	// iota 常量计数器  在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
	const(
		n4 = iota //0
		n5 //1
		n6 //2
		n7 //3
	)
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n6)
	fmt.Println(n7)
}