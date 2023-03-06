package main

import "fmt"

func main() {
	var n int
	fmt.Println("请输入金字塔层数：")
	fmt.Scanln(&n)
	jzt(n)

}
func jzt(n int) {
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
