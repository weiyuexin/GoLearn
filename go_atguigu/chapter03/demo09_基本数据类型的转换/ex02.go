package main

import "fmt"

func main() {
	var n1 int32 = 12
	var n2 int8
	var n3 int8
	n2 = int8(n1) + 127 //编译通过，但是运行结果溢出，结果不是127+12
	//n3 = int8(n1) + 128 //编译出错，因为128已经超出int8的范围了
	fmt.Println("n2=", n2, "n3=", n3) //n2= -117 n3= 0

}
