package main

import "fmt"

func main() {
	// 函数练习
	// 交换两个数的值
	n1, n2 := 3, 5
	swap(&n1, &n2)
	fmt.Println("n1=", n1, "n2=", n2)
}
func swap(n1, n2 *int) {
	tmp := *n1
	*n1 = *n2
	*n2 = tmp
}
