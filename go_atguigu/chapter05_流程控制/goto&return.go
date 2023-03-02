package main

import "fmt"

func main() {
	// TODO goto
	// 1、goto语句可以无条件的转移到程序中指定的行、
	// 2、goto语句通常和条件语句配合使用，可以实现条件转移，跳出循环等
	// 3、不建议使用goto，容易造成流程混乱，维护困难
	// 基本语法：
	/**
	goto label
	...
	label : 语句块
	*/
	//演示goto
	var n int = 10
	fmt.Println("ok1")
	if n > 3 {
		goto label
	}
	fmt.Println("ok2")
	fmt.Println("ok3")
label:
	fmt.Println("ok4")
	fmt.Println("ok5")

	// TODO return
	// return在函数或方法中使用，表示跳出所在函数或方法
	fmt.Println("ok6")
	fmt.Println("ok7")
	fmt.Println("ok8")
	if n > 4 {
		return
	}
	fmt.Println("ok9")
	fmt.Println("ok10")
	fmt.Println("ok11")

}
