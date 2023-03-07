package main

import "fmt"

func main() {
	users := make(map[string]map[string]string)
	modifyUser(users, "tom")
	modifyUser(users, "marry")
	fmt.Println(users)
}
func modifyUser(users map[string]map[string]string, name string) {
	// 判断users中是否有name的key
	//v,ok:=users[name]
	if users[name] != nil { //有这个用户
		users[name]["pwd"] = "88888888"
	} else {
		//没有这个用户
		users[name] = make(map[string]string)
		users[name]["pwd"] = "88888888"
		users[name]["nickname"] = "昵称~" + name
	}
}
