package main

import "fmt"

// 演示go中+的使用
func main() {
	var i = 2
	var j = 3
	var r = i + j
	fmt.Println("r=", r) //r=5

	var str1 = "hello "
	var str2 = "world"
	var res = str1 + str2
	fmt.Println(res)
}
