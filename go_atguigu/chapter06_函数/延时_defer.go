package main

import "fmt"

// todo defer 延时机制 整个函数执行完毕后才会执行defer栈中的语句
func main() {
	//为了在函数执行完成后，及时的释放资源
	//defer最主要的价值就是，当函数执行完毕后，可以及时的释放函数创建的资源
	//如下：模拟代码
	/*func test(){
		//关闭文件资源
		file = openfile(filename)
		defer file.close()
		// 后面还可以使用file
	}*/
	/*func test(){
		//释放数据库资源
		connect = openDatabase()
		defer connect.close()
		// 后面还可以使用connect
	}*/

	res := sum1(10, 20)
	fmt.Println("res=", res) //4
}
func sum1(n1, n2 int) int {
	// 当执行到defer时，会将defer后面的语句压入独立的栈中，暂时不执行
	// 当函数执行完毕后，再从defer栈中，按照先入后出的方式出栈执行
	// 在defer将语句放入到栈时，也会将相关的值拷贝，同时入栈，也就是入栈时值为多少，出栈执行语句时，值还是多少，不会被后面的语句改变

	//执行顺序：
	defer fmt.Println("ok1 n1=", n1) //3
	defer fmt.Println("ok2 n2=", n2) //2
	//增加一句话
	n1++
	n2++

	res := n1 + n2
	fmt.Println("ok3 res=", res) //1
	return res
	/**
	执行结果：
	ok3 res= 30
	ok2 n2= 20
	ok1 n1= 10
	res= 30
	*/
	/**
	增加n1++,n2++后的执行结果：
	ok3 res= 32
	ok2 n2= 20
	ok1 n1= 10
	res= 32
	*/
}
