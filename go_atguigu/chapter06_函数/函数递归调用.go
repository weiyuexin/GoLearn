package main

import (
	"fmt"
)

type myInt int                    //给int取别名，在go中myInt和int虽然都是int类型，但是go认为他们还是两种不同的数据类型
type myFunType func(int, int) int //myFun就是func(int, int)类型的

// todo 递归调用
func main() {
	test1(7)
	var ans = fibo(5)
	fmt.Println(ans)

	fmt.Println(test2(5))  //63
	fmt.Println(monkey(1)) // 1534

	//使用指针修改函数外的变量
	num := 20
	test03(&num)
	fmt.Println("main num=", num)

	//在GO中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了，通过该变量可以对函数调用
	a := getSum1
	fmt.Printf("a的类型是：%T，getSum1的类型是：%T\n", a, getSum1)
	res := a(10, 30)
	fmt.Println("res=", res)

	res2 := myFun(getSum1, 39, 34)
	fmt.Println(res2)

	var num1 myInt
	num1 = 40
	fmt.Println("num1=", num1)
	fmt.Printf("%T\n", num1)
	fmt.Printf("%T\n", num)

	res3 := myFun1(getSum1, 23, 39)
	fmt.Println("res3=", res3)

	// 可变参数
	res4 := sum(6, 3, 4, 5, 6)
	fmt.Println("res4=", res4)

}
func test1(n int) {
	if n > 2 {
		n--
		test1(n)
	} else {
		fmt.Println("n=", n)
	}

	/** 函数递归时需要遵循的原则：
	1) 执行一个函数时，就创建一个新的受保护的独立空间(新函数栈)
	2) 函数的局部变量是独立的，不会相互影响
	3) 递归必须向退出递归的条件逼近，否则就是无限递归:
	4) 当一个函数执行完毕，或者遇到return，就会返回，遵守谁调用，就将结果返回给谁。
	函数执行完毕或者返回时，该函数本身也会被系统销毁
	*/

	// 函数使用的细节和注意事项
	/**
	1、函数形参可以是多个，返回值列表也可以是多个
	2、形参列表和返回值列表的数据类型可以是值类型和引用类型
	3、函数命名遵循标识符命名规则
	4、函数的变量是局部的，函数外不能生效
	5、基本数据类型和 数组 默认都是值传递，即进行值拷贝。在函数内修改，不影响原来的值
	6、如果希望函数内的变量能够修改函数外的变量，可以传入变量的地址&，函数内以指针的方式操作变量
	7、GO不支持函数重载
	8、在GO中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了，通过该变量可以对函数调用
	9、函数既然可以是一种数据类型，因此在go中，函数可以作为形参，并且调用
	10、为了简化数据类型定义，go支持自定义数据类型
		基本语法： type 自定义数据类型名 数据类型  //相当于一个别名
		type mySum func(int,int) int
	11、支持对函数返回值命名
	func cal(n1, n2 int) (sum, sub int) {
		sum = n1 + n2
		sub = n1 - n2
		return
	}
	12、使用_标识符忽略返回值
	13、GO支持可变参数
	// 支持0到多个参数
	func de(args ...int) int {

	}
	//支持1到多个参数,args要写到最后
	func de(n1 int,args ...int) int {

	}
	说明：args是切片，通过args[index]可以访问每个值，可变参数需要放到参数列表最后面
	*/

}

// 13、可变参数
func sum(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

//9、在GO中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了，通过该变量可以对函数调用
func getSum1(n1 int, n2 int) int {
	return n1 + n2
}

//函数既然是一种数据类型，因此在go中，函数可以作为形参，并且调用
func myFun(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}
func myFun1(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

//支持对函数返回值命名
func cal2(n1, n2 int) (sum, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

// 使用指针修改函数外变量
func test03(n *int) {
	*n = *n + 10
	fmt.Println("test03 指针 n=", *n)
}

//练习

//1、斐波那契数列
func fibo(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)

}

// 2、求函数值
func test2(n int) int {
	if n == 1 {
		return 3
	}
	return test2(n-1)*2 + 1
}

// 2、猴子吃桃
func monkey(n int) int {
	if n > 10 || n < 1 {
		fmt.Println("输入的天数不对")
	}
	if n == 10 {
		return 1
	}
	return (monkey(n+1) + 1) * 2
}
