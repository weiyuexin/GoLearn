package main

import (
	"errors"
	"fmt"
)

// todo 自定义错误
func main() {

	// 使用errors.New和panic内置函数自定义错误
	// 1、errors.New("错误说明")，会返回一个errors类型的值，表示一个错误
	// 2、panic内置函数，接收一个interface{}类型的值（也就是任意值）作为参数，
	// 可以接收error类型的变量，输出错误信息并退出程序
	test02()
}

//一个函数，可以读取文件init.conf的信息
func readConf(name string) (err error) {
	if name == "config.ini" {
		//读取
		return nil
	} else {
		//返回自定义错误
		return errors.New("读取文件错误...")
	}
}
func test02() {
	err := readConf("config1.ini")
	if err != nil {
		// 如果读取文件发生错误，就输出这个错误，并终止程序
		panic(err)
	}
	fmt.Println("test02()继续执行")
}
