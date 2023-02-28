package main

import "fmt"

// 演示go中整数类型的使用
func main() {
	// 整数类型
	/**
	类型    有无符号	 占用存储空间
	int8	有			1字节
	int16	有			2字节
	int32	有			4字节
	int64   有			8字节
	*/
	var i int = 1
	fmt.Println("i=", i)

	//演示int8的范围
	var j int8 = -129 //The value of '-129' (type untyped int) cannot be represented by the type int8

}
