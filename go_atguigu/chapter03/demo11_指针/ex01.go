package main

import "fmt"

func main() {
	var num int = 9
	fmt.Println("num address:", &num) //num address: 0xc0000aa058

	var ptr *int
	ptr = &num
	*ptr = 10                //这里修改时，会导致num的值变化
	fmt.Println("num=", num) //num= 10
}
