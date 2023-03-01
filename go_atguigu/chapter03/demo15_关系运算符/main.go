package main

import "fmt"

//关系运算符
func main() {
	//关系运算符结果是bool类型，经常用于if结构和循环结构
	// == != < > <= >=
	var a = 9
	var b = 8
	fmt.Println(a > b)  //true
	fmt.Println(a >= b) //true
	fmt.Println(a < b)  //false
	fmt.Println(a <= b) //false
	fmt.Println(a == b) //false
	fmt.Println(a != b) //true
	flag := a > b
	fmt.Println(flag) //true
}
