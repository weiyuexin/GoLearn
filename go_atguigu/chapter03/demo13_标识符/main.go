package main

import "fmt"

func main() {
	// Golang中严格区分大小写
	var num int = 10
	var Num int = 20
	fmt.Printf("num=%v Num=%v", num, Num) //num=10 Num=20

	// 标识符不能包含空格
	//var ab c int = 30

	// _是一个特殊的标识符，称为空标识符，可以代表任何其他标识符，但是它对应的值会被忽略，所以仅能作为占位符使用
	//var _ int = 23
	//fmt.Println(_) //Cannot use '_' as a value

	// 不能使用保留关键字作为标识符
	//var if int = 10//Type 'int' is not an expression
}
