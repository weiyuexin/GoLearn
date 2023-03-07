package main

import "fmt"

func main() {
	//声明切片
	var slice = make([]int, 0)
	fmt.Println(slice)
	var n int
	fmt.Println("请输入n:")
	fmt.Scanln(&n)
	for i := 0; i <= n; i++ {
		slice = append(slice, fb(i))
	}
	fmt.Println(slice)
}
func fb(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return fb(n-1) + fb(n-2)
}
