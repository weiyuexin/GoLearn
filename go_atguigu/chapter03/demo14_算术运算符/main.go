package main

import "fmt"

// TODO 算术运算符
func main() {
	// + - * / % ++ --

	//除法运算：如果运算的数都是整数，那么除后保留整数部分，去掉小数部分
	fmt.Printf("%d\n", 10/4) //2
	var n1 float32 = 10 / 4
	fmt.Println("n1=", n1) //2
	// 如果我们希望保留小数部分，则需要浮点数参与运算
	var n2 float32 = 10.0 / 4
	fmt.Println("n2=", n2) //n2= 2.5

	//取模 %
	// 公式：a % b = a - a / b * b
	fmt.Println(10 % 3)   //1
	fmt.Println(-10 % 3)  //-1
	fmt.Println(10 % -3)  //1
	fmt.Println(-10 % -3) //-1

	// ++ --
	var i int = 10
	i++
	fmt.Println("i=")
	i--
	fmt.Println("i=", i)

	// 注意事项
	// 1、对于除法/，它的整数除和小数除是有区别的，整数之间做除法时，只保留整数部分
	// 2、当对一个数取模时，可以等价于a%b=a-a/b*b
	// 3、Go中自增自减运算只能当作独立语句使用，不能赋值给变量，也不能在if条件中使用
	// 4、++和--只能写在变量后边，不能写在变量前面

	//var a int
	//a = i++ // ×
	/*if i++ >0{ ×
	}*/
	//++i go中没有前++或前--
}
