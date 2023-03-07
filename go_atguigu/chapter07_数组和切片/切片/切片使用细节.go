package main

import "fmt"

//todo 切片使用细节
func main() {
	//1、切片初始化时，var slice = arr[startIndex,endIndex]
	// 		说明：从var数组下标为startIndex，取到endIndex的元素（不含endIndex）
	//2、切片初始化时，任然不能越界，范围在[0,len(arr)]之间，但是可以动态增长
	/**
	1) var slice = arr[0:end] 可以简写成 var slice = arr[:end]
	2) var slice = arr[start:len(arr)]可以简写成 var slice = arr[start:]
	3) var slice = arr[0:len(arr)] 可以简写成 var slice = arr[:]
	*/
	//3、cap是一个内置函数，用于统计切片的容量，即最大存放元素数量
	//4、切片定义完成后，还不能使用，因为本身是一个空的，需要让其引用一个数组或者make一个空间供切片使用
	//5、切片可以继续切片
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice := arr[1:]
	slice1 := slice[2:4]
	slice1[1] = 100
	// slice和slice1指向的是同一块数据空间，所有改变slice1，其他的都变化
	fmt.Println(slice)  //[20 30 40 100]
	fmt.Println(slice1) //[40 100]

	//append内置函数：用于切片扩容
	var slice3 []int = []int{100, 200, 300}
	fmt.Println("slice3=", slice3) //slice3= [100 200 300]
	// 通过append给slice3追加具体的元素,追加后返回新的切片，不会改变原来的切片
	slice4 := append(slice3, 400, 500, 600)
	fmt.Println("slice3=", slice3) //slice3= [100 200 300]
	fmt.Println("slice4=", slice4) //slice4= [100 200 300 400 500 600]
	//通过append将切片追加给slice3
	slice3 = append(slice3, slice3...)
	fmt.Println("slice3=", slice3) //slice3= [100 200 300 100 200 300]

	//切片append操作的底层原理：
	//1、切片append操作的本质就是对数组的扩容
	//2、go底层会创建一个新的数组newArr（安装扩容后大小）
	//3、将slice原来包含的元素拷贝到新的数组newArr
	//4、slice标识符重新引用到newArr，然后slice原先指向的内存区域会被销毁
	//5、newArr是在底层维护的，程序员是看不到的

	//copy内置函数：切片的拷贝操作
	//copy(slice1,slice2) //将slice2拷贝给slice1，slice2和slice1的数据空间是独立的
	var a []int = []int{1, 2, 3, 4, 5}
	var slice5 []int = make([]int, 10) //切片使用前必须引用数组或者使用make初始化
	copy(slice5, a)
	fmt.Println(slice5) //[1 2 3 4 5 0 0 0 0 0]

	//对上面代码的说明：
	//1、在进行拷贝是，要求两个参数都是切片
	//2、按照上面的代码来看，slice1和slice2的数据空间是独立的，也就是相互不影响
	//3、拷贝时slice1的len小于slice2时，也可以运行，最后slice1输出的是slice2的前几个元素
	var b []int = []int{1, 2, 3, 4, 5}
	var slice6 = make([]int, 1)
	fmt.Println(slice6) //[0]
	copy(slice6, b)
	fmt.Println(slice6) //[1]

	// slice是引用类型，作为函数形参时，在函数中被修改，函数外也会改变
	var c []int = []int{1, 2, 3, 4, 5}
	//fmt.Println(c)/
	test(c)        //[1 2 3 4 5]
	fmt.Println(c) //[100 2 3 4 5]
}
func test(n []int) {
	n[0] = 100
}
