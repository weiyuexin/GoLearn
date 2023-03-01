package main

import "fmt"

func main() {
	// 1、
	var x, y int32 = 30, 34
	if x+y >= 50 {
		fmt.Println("hello world")
	}

	// 2、
	var n1, n2 float64 = 11.0, 18.0
	if n1 > 10.0 && n2 < 20.0 {
		fmt.Println("n1+n2=", n2+n1)
	}

	//3、
	var n3, n4 int32 = 13, 2
	if (n3+n4)%3 == 0 && (n3+n4)%5 == 0 {
		fmt.Println("可以同时被3和5整除")
	} else {
		fmt.Println("不可以同时被3和5整除")
	}

	//4、判断闰年
	var year int = 2021
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Println(year, "是闰年")
	} else {
		fmt.Println(year, "不是闰年")
	}
}
