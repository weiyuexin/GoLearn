package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//1、
	var alphabet [26]byte
	alphabet[0] = 'A'
	for i := 1; i < 26; i++ {
		alphabet[i] = alphabet[i-1] + 1
	}
	for i := 0; i < 26; i++ {
		fmt.Printf("%c ", alphabet[i])
	}
	fmt.Println()

	//2、求出数组最大值和对应下标
	var nums [5]int = [5]int{2, 4, 78, 4, 90}
	max := nums[0]
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			k = i
		}
	}
	fmt.Println("最大值为", max, "对应下标为", k)

	//3、求数组和和平均值
	var arr = [5]int{2, 3, 3, 5, 6}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Printf("和为%v 平均值为%v", sum, float64(sum)/float64(len(arr)))
	fmt.Println()

	// 4、数组反转：随机生成5个数，并将其反转打印
	var intArr [5]int
	//为了每次生成的随机数不一样，我们需要设置一个seed
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(intArr); i++ {
		intArr[i] = rand.Intn(100)
	}
	fmt.Println(intArr)
	// 反转
	for i := 0; i < len(intArr)/2; i++ {
		tmp := intArr[i]
		intArr[i] = intArr[len(intArr)-i-1]
		intArr[len(intArr)-i-1] = tmp
	}
	fmt.Println(intArr)
}
