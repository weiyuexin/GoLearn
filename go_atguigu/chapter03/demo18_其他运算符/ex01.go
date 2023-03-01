package main

import "fmt"

func main() {
	// 求两个数的最大值
	var n1, n2 int = 10, 20
	var max int
	if n1 > n2 {
		max = n1
	} else {
		max = n2
	}
	fmt.Println("max=", max) //max= 20

	//求三个数的最大值
	var n3 int = 50
	if n3 > max {
		max = n3
	}
	fmt.Println("max=", max) //max= 50

}
