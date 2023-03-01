package main

import "fmt"

func main() {
	//switch练习

	//1、
	/*var char byte
	fmt.Println("请输入小写字母：")
	fmt.Scanf("%c", &char)
	switch char {
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
	case 'c':
		fmt.Println("C")
	case 'd':
		fmt.Println("D")
	case 'e':
		fmt.Println("E")
	default:
		fmt.Println("others")

	}*/

	//2、
	/*var score float64
	fmt.Println("请输入学生成绩：")
	fmt.Scanln(&score)
	switch int(score / 60) {
	case 1:
		fmt.Println("合格")
	case 0:
		fmt.Println("不合格")
	default:
		fmt.Println("输入错误")
	}*/

	//3、
	var month byte
	fmt.Println("请输入月份：")
	fmt.Scanln(&month)
	switch month {
	case 3, 4, 5:
		fmt.Println("春季")
	case 6, 7, 8:
		fmt.Println("夏季")
	case 9, 10, 11:
		fmt.Println("秋季")
	case 12, 1, 2:
		fmt.Println("冬季")
	default:
		fmt.Println("输入错误")
	}
}
