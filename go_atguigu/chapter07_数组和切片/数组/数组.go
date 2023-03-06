package main

import "fmt"

// todo 数组
func main() {
	// 定义：数组是可以存放多个同一类型数据的数据类型，
	//数组也是一种数据类型，在GO中数组是值类型，与其他语言不同

	//定义数组
	var hen [6]float64
	//初始化数组
	hen[0] = 3.0
	hen[1] = 5.0
	hen[2] = 1.0
	hen[3] = 3.4
	hen[4] = 2.0
	hen[5] = 50.0
	sum := 0.0
	for _, he := range hen {
		sum += he
	}
	fmt.Printf("totalWeight=%v avg=%v", sum, sum/float64(len(hen)))

}
