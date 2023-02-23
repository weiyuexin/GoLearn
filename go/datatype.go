package main

import {
	"fmt"
	"math"
}	

func main() {
	// 十进制
	var a int =10

	fmt.Printf("%d \n",a) //10 十进制
	fmt.Printf("%b \n",a) //1010 二进制

	// 八进制 以0开头
	var b int = 077
	fmt.Printf("%o \n",b) //77

	// 十六进制 以0x开头
	var c int = 0xff
	fmt.Printf("%x \n",c) //ff
	fmt.Printf("%X \n",c) //FF

	// 浮点型
	fmt.Printf("%f\n", math.Pi)
    fmt.Printf("%.2f\n", math.Pi)


}