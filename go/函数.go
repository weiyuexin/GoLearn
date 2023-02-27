package main

import (
	"errors"
	"fmt"
)

func main() {
	/**
	func 函数名(参数)(返回值){
	    函数体
	}
	函数名：字母、数字、下划线，不能由数字开头，同一包内函数名不能重复
	参数：由参数变量和参数类型构成，多个参数使用,隔开
	返回值：由返回值变量及其变量类型构成，多个返回值使用()包裹，并使用,分隔
	*/
	sayHello()
	ret := intSum(2, 5)
	fmt.Println(ret)
	// 可变参数
	ret1 := intSum2()
	ret2 := intSum2(10)
	ret3 := intSum2(10, 20)
	ret4 := intSum2(10, 20, 30)
	fmt.Println(ret1, ret2, ret3, ret4)
	ret5 := intSum3(100)
	ret6 := intSum3(100, 10)
	ret7 := intSum3(100, 10, 20)
	ret8 := intSum3(100, 10, 20, 30)
	fmt.Println(ret5, ret6, ret7, ret8)
	//多返回值
	fmt.Println(calc(5, 3))
	fmt.Println(calc1(9, 3))
	// 全局变量
	testGlobalVar() //num=10
	// 局部变量
	testLocalVar()
	//fmt.Println(x) //无法访问函数内变量
	testNum() //num=100
	testLocalVar2(3, 4)

	// 定义函数类型
	//add和sub都能赋值给calculation类型的变量
	var c1 calculation
	c1 = add
	fmt.Printf("type of c1:%T\n", c1)
	fmt.Println(c1(2, 3))
	var c2 calculation
	c2 = sub
	fmt.Printf("type of c2:%T\n", c2)
	fmt.Println(c2(5, 2))

	// 函数作为参数
	ret9 := calc2(10, 20, add)
	fmt.Println(ret9)
	// 函数作为返回值
	fmt.Println(do("+"))

	//匿名函数
	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20)
	//自执行函数：匿名函数定义完加()之间执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}
func intSum(x int, y int) int {
	return x + y
}

// 类型简写：相邻的参数变量类型如果一样，那么可以省略类型
func intSum1(x, y int) int {
	return x + y
}
func sayHello() {
	fmt.Println("Hello")
}

// 可变参数：函数的参数数量不固定，使用...来标记,本质上函数的可变参数是通过切片来实现的。
// 注：可变参数通常作为函数最后一个参数
func intSum2(x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

// 固定参数搭配可变参数时，可变参数要放到固定参数后面
func intSum3(x int, y ...int) int {
	fmt.Println(x, y)
	sum := 0
	for _, v := range y {
		sum += v
	}
	return sum
}

// 多返回值
// GO语言支持多返回值，多个返回值必须使用()将所以返回值包裹
func calc(x, y int) (int, int) {
	return x + y, x - y
}

//返回值命名
func calc1(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

//    变量作用域
// 全局变量:全局变量是定义在函数外部的变量，它在程序整个运行周期内都有效。 在函数中可以访问到全局变量。
// 定义全局变量
var num int64 = 10

func testGlobalVar() {
	fmt.Printf("num=%d\n", num)
}

// 局部变量
// 		函数内定义的变量：无法在该函数外使用
//		语句块中定义的变量：if、for、switch等，只在语句块中生效
// 注意：如果局部变量和全局变量重名，优先访问局部变量

//函数内全局变量
func testLocalVar() {
	// 定义一个函数内局部变量，只在该函数内生效
	var x int64 = 100
	fmt.Printf("x=%d\n\n", x)
}

// 测试局部变量和全局变量重名
func testNum() {
	num := 100
	fmt.Printf("num=%d\n", num) //函数中优先使用局部变量
}
func testLocalVar2(x, y int) {
	fmt.Println(x, y) //函数的参数只能在函数内生效
	if x > 0 {
		z := 100 //变量z只在if语句块中生效
		fmt.Println(z)
	}
	// 语句块外不能使用变量z
	//fmt.Println(z)
}

//--------函数类型与变量----------
// 可以使用type关键字定义一个函数类型，格式如下:
type calculation func(int, int) int

//上面语句定义了一个calculation类型，它是一种函数类型，这种函数接收两个int类型的参数并且返回一个int类型的返回值。
//简单来说，凡是满足这个条件的函数都是calculation类型的函数，例如下面的add和sub是calculation类型。
func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}

// ----------高阶函数------------
// 高阶函数分为：
// 1、函数作为参数
// 2、函数作为返回值
// 函数作为参数
func calc2(x, y int, op func(int, int) int) int {
	return op(x, y)
}

// 函数作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

//------------匿名函数和闭包---------------
// 匿名函数：没有函数名的函数
/**匿名函数定义格式：
func(参数列表)(返回值列表){
	函数体
}
注意：匿名函数因为没有函数名，所以没办法像普通函数那样调用，所以匿名函数必须
保存到某个变量或者作为立即执行函数
匿名函数多用于实现回调函数和闭包
*/

// 闭包
// 闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境
