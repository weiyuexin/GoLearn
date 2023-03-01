package main

import "fmt"

func main() {
	//TODO 基本数据类型的默认值
	/**
	类型			默认值
	int			0
	float		0
	string		""
	bool		false
	*/
	var a int       //0
	var b float32   //0
	var c float64   //0
	var d bool      //false
	var name string //""
	// %v 表示按照数据的原值输出
	fmt.Printf("a=%d b=%v c=%v d=%v name=%v", a, b, c, d, name) //a=0 b=0 c=0 d=false name=
}
