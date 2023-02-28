package demo01

import "fmt"

// 定义全局变量
var n10 = 100
var n12 = 200
var name1 = "jack"

// HeroName 定义公有的全局变量
var HeroName string = "宋江"

// 上面的声明改成一次性声明
var (
	n11   = 111
	n13   = 999
	name2 = "www"
)

func main() {
	// 声明/定义变量
	var i int //默认值为0
	// 给变量i赋值
	i = 10
	// 使用变量
	fmt.Println("i=", i) //i= 10
	/**
	变量使用的注意事项：
	1、变量表示内存中的一个存储空间
	2、该区域有自己的名称（变量名）和类型（数据类型）
	3、变量使用的三种方式
		1)指定变量类型，声明后若不赋值，使用默认值
		2)根据变量值自动判断变量类型（类型推导）
		3)省略var，注意：:=左侧的变量不能是已经声明的，否则编译报错
	4、多变量声明
	5、该区域数据值可以在同一个类型范围内变化
	6、变量在同一作用域内不能重名
	7、变量 = 变量名+变量值+数据类型
	8、变量如果没有赋初始值，编译器会使用默认值,int的默认值为0

	*/

	// 变量使用方式
	// 1)不赋值，使用默认值 int的默认值为0
	var a int
	fmt.Println("a=", a) //a=0
	// 2)类型推导 var 变量名 = 变量值
	var b = 2.3
	fmt.Println("b=", b)
	// 3)省略var 变量名 := 变量值
	// b := 333 //No new variables on the left side of ':=' :=左侧不能是已经声明的变量
	c := 78
	fmt.Println("c=", c)

	// 多变量声明
	// 方式一
	var n1, n2, n3 int
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)
	// 方式二
	var n4, n5, n6 = 100, "tom", 90
	fmt.Println("n4=", n4, "n5=", n5, "n6=", n6)
	// 方式三 使用类型推导
	n7, n8, n9 := 2, "wyx", 9
	fmt.Println("n7=", n7, "n8=", n8, "n9=", n9)
	// 一次性声明多个全局变量（函数外声明的变量）
	var (
		name = "weiyeuxin"
		age  = 23
	)
	fmt.Println("name:", name)
	fmt.Println("age:", age)

	// 该区域数据值可以在同一个类型范围内变化
	var i2 = 10
	i2 = 11
	i2 = 23
	//i2 = "ww" //'"ww"' (type string) cannot be represented by the type int
	fmt.Println("i2=", i2)

	// 变量在同一作用域内不能重名
	//var i2 int //'i2' redeclared in this block 作用域内重复定义
}
