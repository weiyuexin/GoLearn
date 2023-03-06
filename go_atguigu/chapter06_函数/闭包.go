package main

import (
	"fmt"
	"strings"
)

// AddUpper 累加器
// 1、返回的是一个匿名函数，但是这个匿名函数引用到函数外的n，
// 因此这个匿名函数就和n形成一个整体，构成闭包、
// 2、可以这样理解：闭包是类，函数是操作，n是字段，函数和它使用到的n构成闭包
// 3、当我们反复调用f函数时，因为n只初始化一次，因此每调用一次就进行累计
// 4、要搞清楚闭包的关键，就是要分析出返回的函数使用了哪些变量，因为函数和他引用到的变量共同构成了闭包

func AddUpper() func(int) int {
	var n int = 20
	var str = "hello"
	return func(x int) int {
		n = n + x
		str = str + "a"
		fmt.Println(str)
		return n
	}
}

// todo 闭包
func main() {
	//定义：一个函数和与其相关的引用环境组合的一个整体就是闭包，引用的变量只会初始化一次，每次引用都会改变变量值

	f := AddUpper()
	fmt.Println(f(1)) //21 执行完成后，n变成了21
	fmt.Println(f(2)) //23 执行完成后，n变成了23
	fmt.Println(f(3)) //26 执行完成后，n变成了26

	//测试makeSuffix是否可以使用
	f1 := makeSuffix(".jpg")
	fmt.Println(f1("www"))
}

//闭包的最佳实践
//func(s string) string {}和suffix构成闭包
// 使用闭包完成的好处：只用传入一次后缀名
func makeSuffix(suffix string) func(string) string {

	return func(s string) string {
		// 如果s没有指定后缀，则加上，否则返回原来的名字
		if !strings.HasSuffix(s, suffix) {
			return s + suffix
		}
		return s
	}
}
