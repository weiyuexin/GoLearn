package main

import "fmt"

// todo string 和 slice
func main() {
	//1、string底层是一个byte数组，因此string也可以进行切片处理
	str := "hello@atguigu"
	// 使用切片获取atguigu
	slice := str[6:]
	fmt.Println("slice=", slice) //slice= atguigu

	//2、string本身是不可变的，也就是说不能通过str[0]='a'的方式来修改字符串
	//str[0] = 'a'//编译不会通过，Cannot assign to str[0]
	//3、如果需要修改字符串，可以先将string->[]byte或[]rune->修改->重新转成string
	// 把hello@atguigu改成zello@atguigu
	arr := []byte(str)
	arr[0] = 'z'
	str = string(arr)
	fmt.Println(str) //zello@atguigu
	// 上面转成[]byte后，可以处理英文和数字，但是不能处理中文
	// 原因是[]byte按照字节处理，而一个汉字是三个字节，因此会出现乱码
	// 解决办法：将string转成[]rune即可，因为[]rune是按字符处理，兼容汉字
	arr1 := []rune(str)
	arr1[0] = '北'
	str = string(arr1)
	fmt.Println(str) //北ello@atguigu
}
