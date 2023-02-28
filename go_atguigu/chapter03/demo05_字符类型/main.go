package main

import (
	"fmt"
)

// 演示字符类型的使用
func main() {
	var c1 byte = 'a'
	var c2 byte = '0' //字符的0
	// 直接输出byte值，就是输出了对应字符的码值
	fmt.Println("c1=", c1) //c1= 97
	fmt.Println("c2=", c2) //c2= 48

	// 如果我们希望输出对应的字符，需要使用到格式化输出
	fmt.Printf("c1=%c c2=%c\n", c1, c2) //c1=a c2=0

	// 如果吧保存的字符在ASCII码表，可以直接保存到byte
	// 如果保存的字符对应的码值大于255，这时，我们考虑使用int类型保存
	// 按照字符的方式输出：fmt.Printf("%c",c1)
	//var c3 byte = '北'
	//fmt.Printf("c3=%c", c3) //cannot use '北' (untyped rune constant 21271) as byte value in variable declaration (overflows)
	var c4 int = '北'
	fmt.Printf("c4=%c", c4)

	// 可以直接给某个变量赋一个数字，然后按照格式化输出%c，会输出该数字对应的unicode字符
	var c5 int = 22269
	fmt.Printf("c5=%c\n", c5) //c5=国

	//字符类型是可以进行运算的，相当于一个整数，运算时按照码值运行
	var n1 = 10 + 'a'
	fmt.Println("n1=", n1)  //n1=107
	fmt.Printf("n1=%c", n1) //n1=k

	//go语言编码默认使用utf-8，再也没有编码的问题
}
