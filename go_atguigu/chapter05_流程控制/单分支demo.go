package main

import "fmt"

// 单分支结构
func main() {
	var age int
	fmt.Println("请输入你的年龄：")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("你的年龄大于18，要对自己的行为负责")
	}
	// golang中支持在if中直接定义一个变量
	if age1 := 20; age1 > 18 { //一条语句时，大括号也不能去
		fmt.Println("你的年龄大于18，要对自己的行为负责")
	}
}
