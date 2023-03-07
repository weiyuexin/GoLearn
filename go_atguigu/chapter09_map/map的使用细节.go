package main

import "fmt"

//todo map的使用细节

// Stu 定义一个学生结构体
type Stu struct {
	Name    string
	Age     int
	Address string
}

func main() {
	//1、map是引用类型，遵守引用类型传递的机制，修改后会直接修改原来的值
	map1 := make(map[int]int, 2)
	map1[1] = 90
	map1[2] = 88
	map1[10] = 1
	map1[20] = 2
	modify(map1)
	fmt.Println(map1) //map[1:90 2:88 10:9999 20:2] 说明map是引用类型
	//2、map的容量达到后，再想map增加元素，会自动扩容，并不会发生panic，也就是说map能动态增长键值对
	//3、map的value也经常使用struct类型，更适合管理复杂的数据
	students := make(map[string]Stu)
	//创建2个学生
	stu1 := Stu{"tom", 18, "北京"}
	stu2 := Stu{"jack", 38, "南京"}
	students["no1"] = stu1
	students["no2"] = stu2
	fmt.Println(students)
	//遍历
	for k, v := range students {
		fmt.Printf("学生的编号:%v\n", k)
		fmt.Printf("学生的姓名:%v\n", v.Name)
		fmt.Printf("学生的年龄:%v\n", v.Age)
		fmt.Printf("学生的地址:%v\n", v.Address)
		fmt.Println("------------------")
	}
}
func modify(map1 map[int]int) {
	map1[10] = 9999
}
