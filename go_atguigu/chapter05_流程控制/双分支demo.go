package main

import "fmt"

// 双分支
func main() {
	/* 基本语法：
	if 条件表达式 {
		语句块1
	} else {
		语句块2
	}
	*/
	var age int
	fmt.Println("请输入你的年龄：")
	fmt.Scanln(&age) //这里一定要记得加&
	if age > 18 {
		fmt.Println("你已经大于18")
	} else {
		fmt.Println("你还没有大于18")
	}
}
