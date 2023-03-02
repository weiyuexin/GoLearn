package main

import "fmt"

// 定义一个函数
func test(n1 int) {
	n1 += 1
	fmt.Println("test n1=", n1)
}

func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	fmt.Println("getSum sum = ", sum)
	return sum
}

//返回多个返回值的函数
func cal1(n1, n2 int) (int, int) {
	return n1 + n2, n1 - n2
}
func main() {
	// todo 函数调用机制
	//   栈区：基本数据类型一般分配在栈区
	//   堆区：引用数据类型一般分配在堆区
	//	 代码区：代码存放在这里

	// 1、在调用一个函数时，会为这个函数分配一个新的栈空间，编译器会通过自身的处理让这个空间和其他的栈空间区分开
	// 2、每一个函数对应的栈中，数据空间是不会混淆的
	// 3、当一个函数调用完毕，程序会销毁这个函数对应的栈空间

	// todo return 语句
	// go函数支持返回多个值
	// 如果返回多个值，在接收时希望忽略某个返回值，则使用_符号表示占位忽略
	// 返回值只有一个，可以不写()

	n1 := 10
	test(n1)
	fmt.Println("main n1=", n1)

	sum := getSum(2, 4)
	fmt.Println("main sum = ", sum)

	// 返回多个返回值
	res1, res2 := cal1(9, 3)
	fmt.Println("sum=", res1, "sub=", res2)
	// 使用_符号忽略返回值
	res3, _ := cal1(12, 3)
	fmt.Println("sum=", res3)

}
