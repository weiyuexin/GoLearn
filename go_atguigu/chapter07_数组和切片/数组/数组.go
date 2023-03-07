package main

import "fmt"

// todo 数组
func main() {
	// 定义：数组是可以存放多个同一类型数据的数据类型，
	//数组也是一种数据类型，在GO中数组是值类型，与其他语言不同

	// 数组定义:
	//var 数组名  [数组大小]数据类型

	//定义数组
	var hen [6]float64
	//初始化数组
	hen[0] = 3.0
	hen[1] = 5.0
	hen[2] = 1.0
	hen[3] = 3.4
	hen[4] = 2.0
	hen[5] = 50.0
	sum := 0.0
	for _, he := range hen {
		sum += he
	}
	fmt.Printf("totalWeight=%v avg=%v\n", sum, sum/float64(len(hen)))

	//数组内存地址
	var intArr [3]int
	//数组默认值为数据类型默认值
	fmt.Println(intArr) //[0 0 0]
	//数组的地址可以通过数组名访问
	//intArr的地址=0xc00000c1b0 intArr[0]的地址=0xc00000c1b0
	fmt.Printf("intArr的地址=%p intArr[0]的地址=%p\n", &intArr, &intArr[0])
	//数组的各个元素的地址间隔是依据数组的类型决定的 int->8 int32->4
	fmt.Printf("intArr[1]的地址=%p intArr[2]的地址=%p\n", &intArr[1], &intArr[2])

	//数组的使用
	/*var score [5]float64
	for i := 0; i < len(score); i++ {
		fmt.Printf("请输入第%d个数", i)
		fmt.Scanln(&score[i])
	}
	for i := 0; i < len(score); i++ {
		fmt.Print(score[i], " ")
	}
	fmt.Println()*/

	// 数组初始化的四种方式
	//1、var nums [3]int = [3]int{1, 2, 3}
	//2、var nums1 = [3]int{1, 2, 3}
	//3、var nums = [...]int{5, 6, 7}
	//4、可以指定元素值对应的下标
	var nums = [3]string{1: "tom", 2: "jack", 0: "marry"}
	fmt.Println(nums) //[marry tom jack]

	//数组的遍历方式for-range
	//1、index是数组下标
	//2、value是在该下标位置的值
	//3、index和value是局部变量
	//4、不想要下标index可以使用_占位符替换
	//5、index和value也可以叫其他名称
	for index, value := range nums {
		fmt.Println("index=", index, "value=", value)
	}
	for _, value := range nums {
		fmt.Println("value=", value)
	}

	//数组的长度是类型的一部分

	// go中数组是值类型，在默认情况下是值传递，因此会进行值拷贝，在函数形参中使用时，数组间不会相互影响
	arr := [3]int{11, 22, 33}
	test(arr)
	fmt.Println("arr in main", arr) //arr in main [11 22 33]
}
func test(arr [3]int) {
	arr[1] = 55
	fmt.Println("arr in test", arr) //arr in test [11 55 33]
}
