package main

import (
	"fmt"
	"unsafe"
)

// 演示bool类型
func main() {
	// 布尔类型：true or false
	// bool类型占1字节
	// bool类型用于流程控制 if or for

	var b bool = false
	fmt.Println("b=", b)
	// bool类型占用1字节
	fmt.Printf("b 占用的空间：%d\n", unsafe.Sizeof(b)) //b 占用的空间：1
}
