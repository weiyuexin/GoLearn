package main

import "fmt"

// TODO 获取用户键盘输入
// 1、fmt.Scanln 按行输入
// 2、fmt.Scanf 格式化输入
func main() {
	// 从控制台接收用户信息【姓名、年龄、薪水、是否通过考试】
	var name string
	var age byte
	var sal float32
	var isPass bool
	// 方式1 ：fmt.Scanln
	fmt.Println("请输入姓名:")
	// 当程序执行到fmt.Scanln时，程序会停止在这里，等待用户输入
	//fmt.Scanln(&name)
	//fmt.Println("请输入年龄:")
	//fmt.Scanln(&age)
	//fmt.Println("请输入薪水:")
	//fmt.Scanln(&sal)
	//fmt.Println("请输入是否通过考试:")
	//fmt.Scanln(&isPass)
	//fmt.Printf("名字：%v \n年龄：%v \n薪水：%v \n是否通过考试：%v", name, age, sal, isPass)

	// 方式2：fmt.Scanf，可以按照指定格式输入
	fmt.Println("请输入姓名、年龄、薪水、是否通过考试，使用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("名字：%v \n年龄：%v \n薪水：%v \n是否通过考试：%v", name, age, sal, isPass)
}
