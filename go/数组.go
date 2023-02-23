package main

import (
	"fmt"
)

func main() {

	// 数组
	// 数组的定义   var 数组变量名 [元素个数]T

	// 定义一个长度为3的int型数组a
	//var a [3]int

	// 数组长度必须是常量，而且长度是数组类型的一部分，一旦定义，数组的长度就不能改变，例如[5]int 和 [10]int是不同的类型，所以不能相互赋值
	//var b [3]int
	//var c [5]int
	// b = c //不能这样赋值

	// 数组的初始化方法

	// 方法一  使用初始化列表来设置元素的值
	var testArray [3]int                        // 数组会初始化为int类型的0
	var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	fmt.Println(testArray)                      //[0 0 0]
	fmt.Println(numArray)                       //[1 2 0]
	fmt.Println(cityArray)                      //[北京 上海 深圳]

	// 方法二  可以让编译器根据初始值的个数自动推断数组长度
	var testArray1 [3]int
	var numArray1 = [...]int{1, 2}
	var cityArray1 = [...]string{"北京", "上海", "深圳"}
	fmt.Println(testArray1)                          //[0 0 0]
	fmt.Println(numArray1)                           //[1 2]
	fmt.Printf("type of numArray:%T\n", numArray1)   //type of numArray:[2]int
	fmt.Println(cityArray1)                          //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n", cityArray1) //type of cityArray:[3]string

	// 方法三 使用索引值的方式初始化数组
	a := [...]int{1: 1, 3: 5}
	fmt.Println(a)                  //[0 1 0 5]
	fmt.Printf("type of a:%T\n", a) //type of a:[4]int

	// 数组遍历的两种方式
	var city = [...]string{"北京", "上海", "深圳"}
	// 方法一
	for i := 0; i < len(city); i++ {
		fmt.Println(a[i])
	}
	//方法二
	for index, value := range city {
		fmt.Println(index, value)
	}

	// 二维数组
	a1 := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a1)       //[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(a1[2][1]) //重庆

	//二位数组遍历
	for _, v1 := range a1 {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
	//注意： 多维数组只有第一层可以使用...来让编译器推导数组长度

	// 数组是值类型，赋值和传参数会复制整个数组，因此改变副本的值，不会改变本身的值

}
