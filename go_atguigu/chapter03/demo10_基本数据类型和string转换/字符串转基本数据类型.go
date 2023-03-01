package main

import (
	"fmt"
	"strconv"
)

// TODO 字符串转基本数据类型
func main() {
	//使用strconv包中的函数

	//字符串转bool
	var str string = "true"
	var b bool
	// strconv.ParseBool()返回两个值，第二个是错误信息error，只想获取bool值，不想获取error，所以使用_忽略掉error
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type : %T b=%v\n", b, b) //b type : bool b=true

	// 字符串转int
	var str2 string = "1234567"
	var n1 int64
	n1, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("n1 type : %T n1=%v\n", n1, n1) //n1 type : int64 n1=1234567

	// 字符串转浮点型
	var str3 string = "2423.24332"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type : %T f1=%v\n", f1, f1) //f1 type : float64 f1=2423.24332

	// 注意事项：
	// 1、string转基本数据类型，需要确保string类型能够转成有效的数据，
	// 比如把“123”转成整数，如果要把“hello”转成整数，就会被直接转成默认值0（其他类型也是一样的）
	var str4 string = "hello"
	var n3 int64
	n3, _ = strconv.ParseInt(str4, 10, 64)
	fmt.Printf("n3 type : %T n3=%v\n", n3, n3) //n3 type : int64 n3=0

}
