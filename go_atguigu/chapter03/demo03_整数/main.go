package main

import (
	"fmt"
	"unsafe" //引入unsafe包，使用Sizeof()函数
)

// 演示go中整数类型的使用
func main() {
	// 整数类型
	/** 有符号
	类型    有无符号	 占用存储空间
	int8	有			1字节
	int16	有			2字节
	int32	有			4字节
	int64   有			8字节
	*/
	var i int = 1
	fmt.Println("i=", i)

	//演示int8的范围 -128 - 127
	//var j int8 = -129 // cannot use -129 (untyped int constant) as int8 value in variable declaration (overflows)

	/**无符号类型
	类型		有无符号		占用存储空间
	uint8	无			1字节
	uint16	无			2字节
	uint32	无			4字节
	uint64	无			8字节
	*/

	// 测试uint8的范围 0-255
	//var k uint8 = -1 //cannot use -1 (untyped int constant) as uint8 value in variable declaration (overflows)
	//fmt.Println("k=", k)

	/** int其他类型
	类型		有无符号		占用存储空间
	int		有			32位系统4字节/64位系统8字节
	uint 	无			同上
	rune	有			同int32，使用Unicode编码
	byte	有			同uint8，存储字符时选用byte
	*/

	// int，uint，byte
	var a int = 8900
	fmt.Println("a=", a)
	var b uint = 1
	var c byte = 255
	fmt.Println("b=", b, "c=", c)

	//整型的使用细节
	var n1 = 100 // 整型的默认声明为int类型
	// 如何查看变量的类型 %T
	fmt.Printf("n1的类型：%T\n", n1)
	// 如何查看一个变量的占用字节大小和数据类型
	var n2 int64 = 10
	//unsafe.Sizeof()时unsafe包的一个函数，可以返回变量占用的字节数
	fmt.Printf("n2的类型：%T  n2占用的字节大小：%d\n", n2, unsafe.Sizeof(n2))

	//变量使用时，遵循保小不保大的原则，在保证程序正确运行下，尽量使用占用空间小的数据类型
	var age byte = 33 // byte 0-255
	fmt.Println("age=", age)
}
