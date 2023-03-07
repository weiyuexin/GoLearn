package main

import "fmt"

//todo map切片
func main() {
	//map切片：切片的数据类型是map
	monsters := make([]map[string]string, 2)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "铁扇公主"
		monsters[1]["age"] = "500"
	}
	//使用append，动态增加切片
	monster := map[string]string{
		"name": "太上老君",
		"age":  "234",
	}
	monsters = append(monsters, monster)
	fmt.Println(monsters)
}
