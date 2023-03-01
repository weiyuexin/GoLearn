package main

import (
	"fmt"
)

func test(byte2 byte) byte {
	return byte2 + 1
}
func main() {
	//基本语法
	/*switch 表达式 {
	case 表达式1,表达式2,...:
		语句块1
	case 表达式3,表达式4,...:
		语句块2
	default:
		语句块
	}*/
	// 注意：
	// 后面golang中case语句块许需要写break，默认就有，当执行玩case语句块后，就直接退出switch

	// 案例
	var key byte
	fmt.Println("请输入一个字符，a,b,c,d,e,f,g")
	fmt.Scanf("%c", &key)
	switch test(key) {
	case 'a':
		fmt.Println("周一")
	case 'b':
		fmt.Println("周二")
	case 'c':
		fmt.Println("周三")
	case 'd':
		fmt.Println("周四")
	case 'e':
		fmt.Println("周五")
	case 'f':
		fmt.Println("周六")
	case 'g':
		fmt.Println("周日")
	default:
		fmt.Println("输入错误")
	}

	// switch 使用注意事项
	/**
	1、switch/case后边是一个表达式（即：常量、变量、一个有返回值的函数等）
	2、case后边的各个表达式的值的数据类型，必须和switch的表达式数据类型一致
	3、case后面可以带多个表达式，使用逗号隔开
	4、case后面表达式如果是常量（字面量），则要求不能重复
	5、case后面不需要带break
	6、default语句不是必须的
	7、switch后面可以不带表达式，类似if-else分支使用
	8、switch后面也可以直接声明一个变量，分号结束，但是不推荐
	9、switch穿透 fallthrough，如果在case语句块后增加fallthrough，则会继续执行下一个case，也叫switch穿透
	10、Type Switch：switch语句还可以被用于type-switch来判断某个interface变量中实际指向的变量类型
	*/

	//switch和if的区别
	// 1、如果判断的是具体的数值，而且数量不多，而且符合整数、浮点数、字符、字符串这几种类型，建议使用switch
	// 2、对区间的判断和结果为bool类型的判断，使用if

	var n1 int32 = 5
	var n2 int32 = 20
	switch n1 {
	case n2, 5, 10: //case后面可以带多个表达式
		fmt.Println("ok")
		//case 5:// 错误 重复常量  Duplicate case 5
		//fmt.Println("ok2")
	default:
		fmt.Println("没有匹配到")
	}

	//switch后面可以不带表达式，相当于if-else
	var age int = 10
	switch {
	case age == 10:
		fmt.Println("age==10")
	case age == 20:
		fmt.Println("age==20")
	}

	//switch也可以对范围进行判断
	var score int = 89
	switch {
	case score > 90:
		fmt.Println("优秀")
	case score > 80:
		fmt.Println("良好")
	case score > 60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}
	//8、switch后面也可以直接声明一个变量，分号结束，但是不推荐

	switch grade := 90; {
	case grade > 90:
		fmt.Println("优秀1")
	case grade > 80:
		fmt.Println("良好1")
	case grade > 60:
		fmt.Println("及格1")
	default:
		fmt.Println("不及格1")
	}

	//switch穿透 fallthrough
	var num int = 10
	switch num {
	case 10:
		fmt.Println("ok10")
		fallthrough //switch穿透
	case 20:
		fmt.Println("ok20")
		fallthrough
	case 30:
		fmt.Println("ok30")
	default:
		fmt.Println("没有匹配到")
	}
}
