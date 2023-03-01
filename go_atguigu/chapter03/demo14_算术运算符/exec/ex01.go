package main

import "fmt"

func main() {
	// 假如还有97天放假，问xx个星期零xx天？
	var days int = 97
	var week int = days / 7
	var day int = days % 7
	fmt.Printf("%d个星期零%d天\n", week, day)

	//定义一个变量保存华氏温度，华氏温度转摄氏温度的公式： 5/9*(华氏温度-100)，求摄氏温度
	var hua float32 = 124.3
	var she float32 = 5.0 / 9 * (hua - 100)
	fmt.Println("摄氏温度:", she)
}
