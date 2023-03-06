package main

import "fmt"

// todo 错误处理机制
func main() {
	test()
	fmt.Println("main()下面的代码")

	// 错误处理的好处：错误处理后，程序不会轻易挂掉，如果加入预警代码，就可以让程序更加健壮

	// 1、默认情况下，当发生错误后(panic)，程序就会退出/崩溃
	// 2、处理方式：defer、panic、recover, recover是一个内置函数，可以捕获异常
	// go中可以抛出一个panic异常，然后在defer中通过recover捕获这个异常，然后正常处理

}
func test() {
	//使用defer+recover来捕获和处理错误
	defer func() {
		err := recover() //recover是一个内置函数，可以捕获异常
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	num1, num2 := 10, 0
	res := num1 / num2
	fmt.Println("res=", res)

}
