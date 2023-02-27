package main

import "fmt"

func main() {
	// 切片
	// 切片（Slice）是一个拥有相同类型元素的可变长度的序列，它是基于数组做的一层封装。
	// 切片是一个引用类型，它内部包含 地址、长度、容量

	// 声明切片的基本语法
	// var name []T
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整数切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	//var d = []bool{false, true}
	fmt.Println(a)        //[]
	fmt.Println(b)        //[]
	fmt.Println(c)        //[false true]
	fmt.Println(a == nil) //true
	fmt.Println(b == nil) //false
	fmt.Println(c == nil) //false
	//fmt.Println(c == d) //切片是引用类型，不支持直接比较，只能和nil比较

	//切片的长度和容量
	// len() 获取切片长度
	// cap() 获取切片容量
	fmt.Println(len(c))
	fmt.Println(cap(c))
	fmt.Println(len(a))
	fmt.Println(cap(a))

	//切片表达式
	a1 := [5]int{1, 2, 3, 4, 5}
	s := a1[1:3] // s := a[low,high]
	fmt.Printf("s:%v  len:%v  cap:%v \n", s, len(s), cap(s))

}
