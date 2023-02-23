package main

import "fmt"

func main() {
	var nums = [...]int{1, 3, 5, 7, 8}
	var ans [10][2]int = getSumNum(nums[:])
	fmt.Println(ans)
}
func getSumNum(nums []int) [10][2]int {
	var ans [10][2]int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == 8 {
				ans[i][0] = i
				ans[i][1] = j
			}
		}
	}
	return ans
}
