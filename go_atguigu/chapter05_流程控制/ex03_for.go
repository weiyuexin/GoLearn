package main

import "fmt"

func main() {
	// 练习一
	// 打印1-100之间所有是9的倍数的整数的个数和总和
	sum := 0
	count := 0
	for i := 1; i < 101; i++ {
		if i%9 == 0 {
			sum += i
			count++
		}
	}
	fmt.Printf("1-100之间所有是9的倍数的整数个数为：%d,总和为：%d\n", count, sum)

	// 练习二
	var num = 6
	for i := 0; i <= num; i++ {
		fmt.Println(i, "\t+\t", num-i, "\t=\t", num)
	}
}
