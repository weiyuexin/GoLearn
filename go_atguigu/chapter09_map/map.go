package main

import "fmt"

// todo map
func main() {
	//基本语法
	//var 变量名 map[keytype]valuetype  key不可以重复，重复的话会覆盖之前的内容
	// 声明是不会分配内存的，初始化需要make，分配内存后才可以赋值和使用
	var a map[string]string = make(map[string]string, 10) //分配10个空间
	//a["no1"] = "宋江"//给空map赋值会报错 assignment to entry in nil map
	a["no1"] = "宋江"
	a["no2"] = "吴用"
	a["no3"] = "武松"
	fmt.Println(a)
	//go中map是无序的

	//注意：
	//1、map在使用前一定要make
	//2、map中key是不能重复的，如果重复了，则以最后的key-value为准
	//3、value是可以相同的
	//4、map中key-value是无序的

	//map使用的三种方式
	//方式1：先声明再make
	var b map[string]string
	b = make(map[string]string, 10)
	b["no1"] = "java"
	//方式2:声明的同时make
	c := make(map[string]string)
	c["no1"] = "c++"
	fmt.Println(c)
	//方式3:在声明的同时直接赋值
	d := map[string]string{"no1": "go", "no2": "c"}
	fmt.Println(d)

	//案例
	stus := make(map[string]map[string]string)
	stus["stu01"] = make(map[string]string, 2)
	stus["stu01"]["name"] = "tom"
	stus["stu01"]["sex"] = "男"
	stus["stu01"]["address"] = "北京"
	stus["stu02"] = make(map[string]string, 2)
	stus["stu02"]["name"] = "marry"
	stus["stu02"]["sex"] = "女"
	stus["stu02"]["address"] = "开封"
	fmt.Println(stus)

	//map的增加和更新
	//map["key"] = value //如果key还没有，就是增加，如果key存在就是修改
	//map删除：delete内置函数
	//delete(map,"key") 如果key存在，就删除key-value，如果key不存在，就不操作，但是也不会报错
	delete(stus, "stu01")
	fmt.Println(stus)

	//map的查找
	//如果stus中存在"stu02",那么就返回true，否则返回false
	val, ok := stus["stu02"]
	if ok {
		fmt.Printf("有no1 key 值为：%v\n", val)
	} else {
		fmt.Println("没有 no1 key")
	}

	//map没有提供一次性删除所有key-value的方法，想要删除所有，有以下两种方法：
	//1、遍历逐一删除
	//2、直接make一个新的空间，原先的就会被回收
	stus = make(map[string]map[string]string)
	fmt.Println(stus) //map[]

}
