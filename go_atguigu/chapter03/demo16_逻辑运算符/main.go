package main

import "fmt"

//测试函数
func test() bool {
	fmt.Println("test")
	return true
}

// TODO 逻辑运算符
func main() {
	// && 逻辑与(短路与)，第一个条件为false，则第二个条件不会判断
	//|| 逻辑或(短路或)，第一个为true，则第二个条件不会再判断
	//! 逻辑非

	// &&
	var age = 40
	var i int = 10
	if age > 30 && age < 50 {
		fmt.Println("ok1")
	}
	if age > 30 && age < 40 {
		fmt.Println("ok2")
	}
	if i < 9 && test() {
		fmt.Println("ok")
	}
	if i > 9 || test() {
		fmt.Println("ok")
	}

	// ||
	if age > 30 || age < 50 {
		fmt.Println("ok3")
	}
	if age > 30 || age < 40 {
		fmt.Println("ok4")
	}

	//!
	if age > 30 {
		fmt.Println("ok5")
	}
	if !(age > 30) {
		fmt.Println("ok2")
	}

}
