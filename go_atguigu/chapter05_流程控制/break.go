package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//我们为了生成一个随机数，还需要为rand设置一个种子
	//time.Now().Unix() : 返回从1970.1.1到现在的秒数
	rand.Seed(time.Now().Unix())
	// 生成随机数
	n := rand.Intn(101) //[0,101)
	fmt.Println(n)

	var count = 0
	for {
		rand.Seed(time.Now().UnixNano()) //纳秒
		n := rand.Intn(101)              //[0,101)
		fmt.Println(n)
		count++
		if n == 99 {
			break
		}
	}
	fmt.Println("生成99一共循环了", count)

	// break使用注意事项
	// 1、break语句出现在多层嵌套循环的语句中，可以通过标签指明要终止的是那一层语句块

	// 演示指定标签形式使用break
lable1:
	for i := 0; i < 4; i++ {
		//lable2:
		for j := 0; j < 10; j++ {
			if j == 2 {
				//break //break会默认跳出最近的for循环
				break lable1
			}
			fmt.Println("j=", j)
		}
	}

	// 练习
	//1、
	sum := 0
	for i := 1; i < 100; i++ {
		sum += i
		if sum > 20 {
			fmt.Println("第一次大于20时的对应数为：", i)
			break
		}
	}

	//2、
	count = 3
	var name string
	var pass string
	for count > 0 {
		fmt.Println("请输入用户名和密码：")
		fmt.Scanln(&name)
		fmt.Scanln(&pass)
		if name == "张无忌" && pass == "888" {
			fmt.Println("登录成功")
			break
		} else {
			count--
			fmt.Printf("用户名或密码错误，还有%d次机会\n", count)
		}
	}
}
