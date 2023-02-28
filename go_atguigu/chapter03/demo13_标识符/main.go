package main

import (
	// 引入包
	"GoLearn/go_atguigu/chapter03/demo01"
	"fmt"
)

func main() {
	// Golang中严格区分大小写
	var num int = 10
	var Num int = 20
	fmt.Printf("num=%v Num=%v\n", num, Num) //num=10 Num=20

	// 标识符不能包含空格
	//var ab c int = 30

	// _是一个特殊的标识符，称为空标识符，可以代表任何其他标识符，但是它对应的值会被忽略，所以仅能作为占位符使用
	//var _ int = 23
	//fmt.Println(_) //Cannot use '_' as a value

	// 不能使用保留关键字作为标识符
	//var if int = 10//Type 'int' is not an expression

	//  标识符命名规范
	// 1、包名：尽量保存package的名字和目录一致，尽量采取有意义的包名，简短，有意义，不要和标准库冲突
	// 2、变量名、函数名、常量名：采用驼峰命名法
	// 3、如果变量名、函数名、常量名首字母大写，则可以被其他的包访问，
	// 如果首字母小写，则不能在其他包中使用（注：可以简单理解成首字
	// 母大写是公有的，首字母小写是私有的，在golang中没有public、private等访问控制关键字）

	// 使用utils.go中的heroName 包名.标识符
	fmt.Println(demo01.HeroName)
}
