package main

import "fmt"

func main() {
	// if else 分支结构
	// if 表达式1 {
	// 	分支1
	// } else if 表达式2 {
	// 	分支2
	// } else {
	// 	分支3
	// }
	score := 80
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	// 特殊写法
	if score1 := 65; score1 >= 90 {
		fmt.Println("A")
	} else if score1 > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	// for循环
	// 常用
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// 不常用，初始语句可以忽略，但是初始化语句后的分号必须要写
	b := 0
	for ; b < 10; b++ {
		fmt.Println(b)
	}
	// 不常用，初始语句和结束语句都可以省略
	c := 0
	for c < 10 {
		fmt.Println(c)
		c++
	}
	// 无限循环
	// for {
	// 	fmt.Println("hhh")
	// }
	for i := 0; i < 10; i++ {
		fmt.Print("循环", i, "次")
		if i == 6 {
			break
		}
	}
	fmt.Println()
	for i := 0; i < 10; i++ {
		if i == 6 {
			continue
		}
		fmt.Print("循环", i, "次")
	}
	fmt.Println()

	// switch case
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小指")
	default:
		fmt.Println("无效输入")
	}
	// 一个分支可以有多个值，使用逗号分隔
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
	//分支还可以使用表达式，这时switch语句后就不用跟判断变量
	age := 30
	switch {
	case age < 12:
		fmt.Println("少年")
	case age < 25:
		fmt.Println("青年")
	case age < 45:
		fmt.Println("中年")
	case age < 60:
		fmt.Println("中老年")
	}
	//fallthrough语法可以执行满足条件的case的下一个case
	s := 'a'
	switch {
	case s == 'a':
		fmt.Println("a")
		fallthrough
	case s == 'b':
		fmt.Println("b")
	case s == 'c':
		fmt.Println("c")
	default:
		fmt.Println("...")
	}

	// goto 调转到指定标签，使用goto可以跳出循环
	// 使用goto简化跳出循环的代码
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				//设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
breakTag:
	fmt.Println("结束for循环")

	// break结束循环，还可以在break语句后添加标签，表示退出某个标签对应的代码块
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 5 {
				break BREAKDEMO1
			}
			fmt.Println("%v-%v\n", i, j)
		}
	}
	fmt.Println("..........")

	// continue 结束当前循环，开始下一次循环迭代
forloop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Println("%v-%v", i, j)
		}
	}

}
