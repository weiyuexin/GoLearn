package main

import "fmt"

// 赋值运算符
func main() {
	// = += -= *= /= %=
	// 位运算符
	// <<= 左移后赋值
	// >>= 右移后赋值
	// &= 按位与后赋值
	// ^= 按位异或后复制
	// |= 按位或后赋值

	// 交换两个变量的值
	a := 9
	b := 7
	// 定义临时变量
	t := a
	a = b
	b = t
	fmt.Printf("交换后：a=%v,b=%v \n", a, b) //交换后：a=7,b=9
	// 复合赋值
	a += 10
	fmt.Println("a=", a) // 7+10=17

	// 面试题：交换两个变量的值，不使用中间变量
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("a=%d,b=%d", a, b) //a=9,b=17
}
