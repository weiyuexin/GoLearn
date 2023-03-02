package main

import "fmt"

func main() {
	// for循环基本语法
	/*for 循环变量初始化; 循环条件; 循环变量迭代 {
		循环体
	}*/

	// for的第二种使用方法
	// 变量初始化和变量迭代写在其他地方
	/*for 循环条件 {
		循环体
	}*/

	// for循环的第三种使用方法，是一个无限循环，一般和break配合使用
	// 也等价于for ;;{}
	/*for{
		循环体
	}*/

	// for-range 方便遍历数组和字符串

	for i := 0; i < 10; i++ {
		fmt.Println("Hello Golang-", i)
	}
	// for循环的第二种写法，类似其他语言的while
	var j = 0    //循环变量初始化
	for j < 10 { //循环条件
		fmt.Println("Hello Golang-", j)
		j++ //循环变量迭代
	}
	//for循环的第三种使用方法，这种写法通常配合break使用
	k := 0
	for {
		fmt.Println("ok")
		if k == 10 {
			break // 跳出for循环
		}
		k++
	}

	// 如何遍历字符串
	// 方式一；传统方式
	var str string = "hello world北京"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}
	// 方式二：for-range
	str = "jabbed_72348&&&* &*北京"
	for index, val := range str {
		fmt.Printf("index=%d,val=%c\n", index, val)
	}
	// 如果字符串含有中文，那么传统的遍历方式就会出错，会出现乱码，
	// 原因  :  传统的遍历方式对字符串的遍历是按照字节来遍历，
	// 而一个汉字在utf-8编码中对应3个字节
	// 如何解决：需要将str转成切片
	// 对应for-range遍历方式，默认是按照字符方式遍历，因此有汉字也是ok的
	str = "hello world北京"
	str2 := []rune(str) //把string转成切片
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c\n", str2[i])
	}

	// go中没有while和do..while，可以使用for循环实现
	// 1、for循环实现while效果
	/*for {
		if 循环条件表达式 {
			break //跳出循环
		}
		循环操作（语句）//可能一次都不执行
		循环遍历迭代
	}*/
	// 2、for循环实现do...while效果
	/*for {
	    循环操作（语句）//至少执行一次
		循环遍历迭代
		if 循环条件表达式 {
			break //跳出循环
		}
	}*/

	// 使用while实现输出十句话
	m := 0
	for {
		if m == 10 {
			break
		}
		fmt.Println("Hello world", m)
		m++
	}
	fmt.Println("m=", m)

	//使用do...while
	var jj = 1
	for {
		fmt.Println("hello ok")
		jj++
		if jj > 10 {
			break
		}
	}
}
