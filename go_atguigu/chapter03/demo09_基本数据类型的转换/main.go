package main

import "fmt"

// TODO 基本数据类型的相互转换
func main() {
	// GO与Java/c不同，不同类型的变量之间赋值时需要显式转换(强制类型转换)，GO中数据类型不能自动转换
	// 基本语法：
	// 表达式 T(v) 表示将值v转换为类型T
	// T：就是数据类型
	// v：需要转换的变量
	var i int = 100
	fmt.Printf("i的类型：%T\n", i) //i的类型：int
	// i => float32
	var n1 float32 = float32(i)
	fmt.Printf("n1的类型：%T\n", n1) //n1的类型：float32
	// i => int8
	var n2 int8 = int8(i)
	fmt.Printf("n2的类型：%T\n", n2) //n2的类型：int8
	fmt.Printf("i的类型：%T\n", i)   //i的类型：int

	// 细节说明
	// 1、GO中，数据类型的转换可以是从表示范围小->表示范围大，也可以是从表示范围大->表示范围小
	// 2、被转换的是变量出处的数据，变量本身的数据类型不会发生变化
	// 3、在转换中，将int64转换为int8，编译不会报错，只是转换的结果会发生溢出

	// 溢出
	var num1 int64 = 999999
	var num2 int8 = int8(num1)
	fmt.Println("num2=", num2) //num2= 63
}
