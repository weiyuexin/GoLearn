package main

import "fmt"

//todo 切片遍历
func main() {
	//方式1 常规for循环
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice := arr[1:4]                 // 20 30 40
	for i := 0; i < len(slice); i++ { //slice[0]=20 slice[1]=30 slice[2]=40
		fmt.Printf("slice[%v]=%v ", i, slice[i])
	}
	fmt.Println()
	//方式2、for-range
	for i, v := range slice { //i=0 v=20  i=1 v=30  i=2 v=40
		fmt.Printf("i=%v v=%v  ", i, v)
	}
}
