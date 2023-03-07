package main

import "fmt"

//todo 切片 slice 个数不确定的数组
func main() {
	//1、切片是数组的一个引用，因此切片是引用类型，在传值时，遵守引用传递的机制
	//2、切片的使用和数组类似，遍历切片、访问切片元素和求切片长度len(slice)都是一样的
	//3、切片的长度可以变化，因此切片相当于一个动态数组
	//4、切片定义的基本语法：
	//var 变量名 []类型
	var a []int
	fmt.Println(a) //[]

	//演示切片的基本使用
	var intArr [5]int = [...]int{1, 22, 33, 66, 88}

	//声明/定义一个切片
	//1、slice：切片名
	//2、intArr[1:3]表示slice引用到intArr这个数组
	//3、引用intArr数组的起始下标为1，终止下标为3,但是不包含3
	slice := intArr[1:3]
	fmt.Println("slice 的元素是 = ", slice)        //[22 33]
	fmt.Println("slice 的元素个数 = ", len(slice))  //slice 的元素个数 =  2
	fmt.Println("slice 的当前容量是 = ", cap(slice)) //slice 的当前容量是 =  4

	fmt.Printf("intArr[1]的地址=%p\n", &intArr[1]) //intArr[1]的地址=0xc0000cc038
	fmt.Printf("slice[0]的地址=%p\n", &slice[0])   //slice[0]的地址=0xc0000cc038

	//总结：
	//1、slice的确是一个引用类型
	//2、从底层来说，其实slice就是一个结构体
	/**
	type slice struct{
		ptr *[2]int //存放数据
		len int //当前存储数据的数量
		cap int //切片的容量
	}
	*/

	//切片使用的三种方式
	//1、定义一个切片，然后让切片去引用一个已经创建好的数组，如上22行：slice := intArr[1:3]

	//2、通过make来创建切片
	//基本语法：var 切片名 []type = make([]type,len,cap)  len:大小 cap:容量，可选
	var slice1 []int = make([]int, 4, 10) //[0 0 0 0]
	fmt.Println(slice1)

	//对于切片，必须make后才可以使用
	var slice2 []float64 = make([]float64, 5, 10)
	fmt.Println(slice2)
	//总结：
	//1)通过make方式创建切片可以指定切片的大小和容量
	//2)如果没有给切片的各个元素赋值，那么就会使用默认值
	//3)通过make方式创建的切片对应的数组是由make底层维护，对外不可见，即只能通过slice去访问各个元素

	//3、定义一个切片，直接就指定具体数组，使用原理类似make的方式
	var slice3 []int = []int{1, 2, 3}
	fmt.Println(slice3) //[1 2 3]

	var strSlice []string = []string{"java", "cpp", "golang"}
	fmt.Println("strSlice=", strSlice)           //strSlice= [java cpp golang]
	fmt.Println("strSlice size=", len(strSlice)) //strSlice size= 3
	fmt.Println("strSlice cap=", cap(strSlice))  //strSlice cap= 3

	//方式一和方式二的区别：
	//1、方式1是直接引用数组，这个数组是事先存在的，程序员可见
	//2、通过make来创建切片，make会创建一个数组，是由切片在底层维护的，程序员不可见

}
