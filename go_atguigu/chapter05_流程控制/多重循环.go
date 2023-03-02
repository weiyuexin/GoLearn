package main

import "fmt"

func main() {
	//1、统计三个班成绩情况，每个班5个学生，求出各个班的平均分和所有班级的平均分
	//2、统计三个班的及格人数
	sum1 := 0.0
	var classNum int = 3
	var stuNum int = 5
	var passCount int = 0
	for j := 1; j <= classNum; j++ {
		sum := 0.0
		for i := 1; i < 6; i++ {
			var score float64
			fmt.Printf("请输入第%d个班级第%d个学生成绩：\n", j, i)
			fmt.Scanln(&score)
			//累计总分
			sum += score
			//判断及格人数
			if score >= 60 {
				passCount++
			}
		}
		fmt.Println(j, "班级的平均分是：", sum/float64(stuNum))
		sum1 += sum
	}
	fmt.Println("所有班级的平均分是：", sum1/(float64(stuNum)*float64(classNum)))
	fmt.Println("及格人数：", passCount)
}
