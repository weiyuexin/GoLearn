package main

import "fmt"

// 演示字符串类型
func main() {
	// 字符串是一串固定长度的字符连接起来的字符序列，GO字符串是由单个字节连接起来。GO的字符串的字节使用UTF-8编码标识Unicode文本
	//string的基本使用
	var address string = "北京海淀 四季春 "
	fmt.Println(address)

	//go中字符串是不可变的
	var str = "hello"
	//str[0] = 'a' //Cannot assign to str[0]  这里就不能去修改str的内容，即go中字符串是不可变的
	fmt.Println(str)

	// 字符串的两种表示形式
	// 1、双引号：会识别转义字符
	// 2、反引号：以字符串原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代码等效果
	var str2 = "abc\nabc"
	fmt.Println(str2) //abc 换行 abc
	var str3 = `abc\nabc`
	fmt.Println(str3) //abc\nabc

	// 字符串拼接方式
	var str4 = "hello " + "world "
	str4 += "哈哈哈"
	fmt.Println(str4)
	//当一个拼接的操作很长时，可以分行写,但是要把+号留到上一行
	var str5 = "hello " + "world" + "hello " +
		"world" + "hello " + "world" + "hello " +
		"world" + "hello " + "world"
	fmt.Println(str5)
}
