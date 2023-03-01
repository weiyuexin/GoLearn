package main

import "fmt"

// TODO 指针类型
func main() {
	// 基本数据类型的内存布局
	var i int = 10
	// 获取变量的地址，用&
	// i的地址,&i
	fmt.Println("i的地址=", &i) //i的地址= 0xc0000120b8

	// 1、ptr是一个指针变量
	// 2、ptr的类型 *int
	// 3、ptr 本身的值 &i
	var ptr *int = &i
	fmt.Printf("ptr的值：%v\n", ptr)   //ptr的值：0xc0000120b8
	fmt.Printf("ptr的地址：%v\n", &ptr) //ptr的地址：0xc0000ce020

	// 获取指针类型所指向的值，使用：* ，比如 var ptr *int ,使用*ptr获取ptr指向的值
	fmt.Printf("ptr指向的值：%v\n", *ptr) //ptr指向的值：10

	//指针的使用细节
	//1、值类型都有对应的指针类型，形式为 *数据类型
	//2、值类型包括：基本数据类型int、float、bool、string、数组和结构体struct

}
