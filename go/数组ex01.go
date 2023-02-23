package main

import "fmt"

func main() {
	var nums = [...]int{1, 3, 5, 7, 8}
	var ans int = sum(nums[:])
	fmt.Println(ans)
}

// 求数组所有元素的和
func sum(nums []int) int {
	var ans int = 0
	for i := 0; i < len(nums); i++ {
		ans += nums[i]
	}
	return ans
}
