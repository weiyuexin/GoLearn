package main

import "fmt"

//演示浮点类型
func main() {
	var price float32 = 89.12
	fmt.Println("price=", price)
	// 小数类型分类
	// 类型		占用存储空间
	//单精度float32 	4字节
	//双精度float64	8字节
	var num1 float32 = -0.00089
	var num2 float64 = -67899087.09
	fmt.Println("num1=", num1, "num2=", num2) //num1= -0.00089 num2= -6.789908709e+07 科学计数法
	// 尾数部分可能会丢失精度
	var num3 float32 = -123.000009001
	var num4 float64 = -123.000009001
	fmt.Println("num3=", num3, "num4=", num4) //num3= -123.00001 num4= -123.000009001 float32单精度造成精度损失
	// 说明：float64的精度比float32的要高
	// 如果希望保存精度高的数，则应该选择float64
	// 浮点数存储部分分为三部分：符号位+指数位+尾数位

	// 使用细节
	// go的浮点类型有固定的范围和字段长度，不受os的影响
	// go的浮点类型默认声明为float64
	// 浮点型常量有两种表示形式
	//		十进制形式：5.12  .234
	//		科学计数法形式： 5.1234e2 = 5.12 * 10的2次方 5.12E-2 = 5.12/10的2次方
	// 通常使用float64

	// 默认类型
	var num5 = 1.1
	fmt.Printf("num5的数据类型: %T\n", num5) //num5的数据类型: float64

	// 十进制形式
	num6 := 5.12
	num7 := .23
	fmt.Println("num6=", num6, "num7=", num7) //num6= 5.12 num7= 0.23

	//科学计数法
	num8 := 5.1234e2
	num9 := 5.1234e2
	fmt.Println("num8=", num8, "num9=", num9) //num8= 512.34 num9= 512.34
	num10 := 5.1234e-2
	fmt.Println("num10=", num10) //num10= 0.051234
}
