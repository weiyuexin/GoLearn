package main

import (
	"fmt"
	"math"
)

// TODO 多分支
func main() {
	/*if 条件表达式1 {
		代码块1
	}else if 条件表达式2 {
		代码块2
	}...
	}else {
		代码块n
	}*/
	var score int
	fmt.Println("请输入岳云鹏的成绩：")
	fmt.Scanln(&score) //又忘记写&
	if score == 100 {
		fmt.Println("BMW一辆")
	} else if score > 80 {
		fmt.Println("iphone 7 Plus")
	} else if score >= 60 {
		fmt.Println("ipad")
	} else {
		fmt.Println("null")
	}
	/*var b bool = true
	if b = false { // 这样写在golang中是无法编译通过的，不允许将赋值语句放到if条件表达式中

	}*/

	// 案例三
	var a, b, c float64 = 3, 100, 6
	var x1, x2 float64
	if b*b-4*a*c < 0 {
		fmt.Println("无解")
	} else if b*b-4*a*c == 0 {
		x1 = (-b + math.Sqrt(b*b-4*a*c)) / 2 * a
		fmt.Println("x=", x1)
	} else {
		x1 = (-b + math.Sqrt(b*b-4*a*c)) / 2 * a
		x2 = (-b - math.Sqrt(b*b-4*a*c)) / 2 * a
		fmt.Println("x1=", x1, "x2=", x2)
	}

	// 案例四
	var height int32
	var money float32
	var handsome bool
	fmt.Println("请输入身高、财富和是否帅：")
	fmt.Scanf("%d %f %t", &height, &money, &handsome)
	if height > 180 && money > 1000 && handsome {
		fmt.Println("我一定要嫁给他")
	} else if height > 180 || money > 1000 || handsome {
		fmt.Println("嫁吧，比上不足，比下有余")
	} else {
		fmt.Println("滚")
	}
}
