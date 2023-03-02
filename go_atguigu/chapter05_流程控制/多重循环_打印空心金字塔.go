package main

import (
	"fmt"
)

func main() {
	// 打印经典金字塔
	/*var n = 0
	fmt.Println("请输入金字塔层数：")
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}*/

	// 打印空心金字塔
	var n = 0
	fmt.Println("请输入金字塔层数：")
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		for j := n - i; j > 0; j-- {
			fmt.Print(" ")
		}
		for j := 0; j < i*2-1; j++ {
			if j == 0 || j == i*2-2 || i == n {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
