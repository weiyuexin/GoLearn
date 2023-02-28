package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 基本数据类型转string
	// 1、fmt.Sprintf("%参数",表达式)  推荐
	// 2、使用strconv包的函数

	var num1 int = 99
	var num2 float64 = 23.213
	var b bool = true
	var char byte = 'h'
	var str string //空的string

	//※使用第一种方式转换 fmt.Sprintf("%参数",表达式)
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type : %T str=%v\n", str, str) //str type : string str=99
	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type : %T str=%v\n", str, str) //str type : string str=23.213000
	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type : %T str=%v\n", str, str) //str type : string str=true
	str = fmt.Sprintf("%c", char)
	fmt.Printf("str type : %T str=%v\n", str, str)

	//使用第二种方式转换，使用strconv函数
	var num3 int = 99
	var num4 float64 = 23.234234
	var b2 bool = false
	//转换十进制
	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T  str=%q\n", str, str) //str type string  str="99"

	// 说明：'f' 格式  10：表示小数位保留10位    64：表示这个小数是float64
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T  str=%q\n", str, str) //str type string  str="23.2342340000"

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T  str=%q\n", str, str) //str type string  str="false"

	// strconv包中还有一个函数：Itoa() : 把int转换成字符串
	var num5 int = 88
	str = strconv.Itoa(num5)
	fmt.Printf("str type %T  str=%q\n", str, str) //str type string  str="88"
}
