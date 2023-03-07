package main

import "fmt"

// todo map遍历
func main() {
	//map遍历使用for-range遍历
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	for k, v := range cities {
		fmt.Printf("k=%v v=%v\n", k, v)
	}

	stus := make(map[string]map[string]string)
	stus["stu01"] = make(map[string]string, 2)
	stus["stu01"]["name"] = "tom"
	stus["stu01"]["sex"] = "男"
	stus["stu01"]["address"] = "北京"
	stus["stu02"] = make(map[string]string, 2)
	stus["stu02"]["name"] = "marry"
	stus["stu02"]["sex"] = "女"
	stus["stu02"]["address"] = "开封"
	for k, v := range stus {
		fmt.Println("k=", k)
		for k1, v1 := range v {
			fmt.Printf("\t k1=%v v1=%v \n", k1, v1)
		}
		fmt.Println()
	}

	//map的长度：len()
	fmt.Println(len(stus))   //2
	fmt.Println(len(cities)) //3
}
