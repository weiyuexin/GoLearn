package main

import "fmt" //格式化输出输入

func main() {
	/** 转义字符
	\t 制表符，实现对齐功能
	\n 换行符
	\\ 表示一个\
	\" 表示一个"
	\r 回车，表示从当前行的最前面开始输出，覆盖掉以前的内容
	*/
	// 演示转义字符的使用
	fmt.Println("tom\tjack")                  //tom     jack
	fmt.Println("tom\njack")                  //tom 换行 jack
	fmt.Println("D:\\尚硅谷GO资料\\代码\\chapter02") //D:\尚硅谷GO资料\代码\chapter02
	fmt.Println("tom:\"I Love You\"")         //tom:"I Love You"
	fmt.Println("哈哈哈哈哈哈\r张飞")                 //张飞哈哈哈哈

}
