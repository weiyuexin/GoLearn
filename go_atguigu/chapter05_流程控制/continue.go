package main

import "fmt"

// TODO continue
func main() {
	//1、continue语句用于结束本次循环，继续执行下一次循环
	//2、continue出现在多层嵌套循环中时，可以通过标签指明要跳过那一层循环
lable1:
	for i := 0; i < 4; i++ {
		//lable2:
		for j := 0; j < 10; j++ {
			if j == 2 {
				//break //break会默认跳出最近的for循环
				continue lable1
			}
			fmt.Println("j=", j)
		}
	}

	//练习
	//1、
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}
	//2、
	var num int
	var positive = 0
	var negative = 0
	for {
		fmt.Println("请输入数字：")
		fmt.Scanln(&num)
		if num == 0 {
			break
		}
		if num > 0 {
			positive++
		} else if num < 0 {
			negative++
		}
	}
	fmt.Println("正数个数：", positive, " 负数个数：", negative)
}
