package main

import "fmt"

func main() {
	//TODO 嵌套分支
	//嵌套分支不宜太多

	//案例一
	/*var score float32
	var sex int32
	fmt.Println("请输入成绩和性别：")
	fmt.Scanf("%f %c", &score, &sex)
	if score < 8 {
		if sex == '男' {
			fmt.Println("进入男子组决赛")
		} else {
			fmt.Println("进入女子组决赛")
		}
	} else {
		fmt.Println("么有进入决赛")
	}*/

	//案例二
	var month, age byte
	fmt.Println("请输入月份和年龄：")
	fmt.Scanf("%d %d", &month, &age)
	if month >= 4 && month <= 10 {
		if age >= 60 {
			fmt.Println("票价：20元")
		} else if age >= 18 {
			fmt.Println("票价：60元")
		} else {
			fmt.Println("票价：30元")
		}
	} else {
		if age >= 18 && age < 60 {
			fmt.Println("票价：40元")
		} else {
			fmt.Println("票价：20元")
		}
	}

}
