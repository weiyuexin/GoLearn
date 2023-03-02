package main

import (
	"fmt"
	"packagedemo/utils" //引入自定义包，项目不在gopath中时，需要使用go mod定义模块
)

// TODO 包
func main() {
	// 包的作用
	// 1、区分相同名字的函数、变量名等标识符
	// 2、当程序文件很多时，可以很好的管理项目
	// 3、控制函数、变量等访问范围，即作用域

	// 包的基本语法
	// package 包名
	// import "包路径"

	n1, n2 := 1.2, 2.3
	// 调用其他包中的函数
	result := utils.Cal(n1, n2, '+')
	fmt.Println(result)
}
