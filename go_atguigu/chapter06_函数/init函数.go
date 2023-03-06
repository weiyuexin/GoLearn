package main

import "fmt"

// 为了看到全局变量是被先初始化的，我们先写一个函数
func initvar() int {
	fmt.Println("变量定义")
	return 90
}

var age = initvar()

//todo init函数
func main() {
	//每一个源文件都可以包含一个init函数，该函数会在main函数执行前，被go运行框架调用，
	//也就是说init会在main函数前被调用，多用于做初始化操作
	fmt.Println("main() age=", 90) //main函数在init函数后被调用

	// init函数使用细节和注意事项
	/**
	1、如果一个文件中包含全局变量定义、init函数、main函数，则执行流程是：
			全局变量定义=》init函数=》main函数
			被引入的包中的init先执行
	2、init函数最主要的作用，就是完成一些初始化工作

	3、被引入的包中的全局变量定义 =》被引入的包中的init函数 =》全局变量定义=》init函数=》main函数

	*/

}
func init() {
	fmt.Println("init()") //init函数先调用
}
