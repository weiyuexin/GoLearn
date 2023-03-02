package main

import "fmt"

func main() {
	//函数的基本语法
	/*func 函数名(形参列表) (返回值列表){
		语句块
		return 返回值列表
	}*/
	//1、形参列表：表示函数的输入
	//2、语句：为了实现某一功能的代码块
	//3、函数可以有返回值，也可以没有，golang中支持多返回值
	result := cal(1.2, 2.3, '/')
	fmt.Println("运算结果:", result)
}

// 封装计算功能
func cal(n1 float64, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	}
	return res
}
