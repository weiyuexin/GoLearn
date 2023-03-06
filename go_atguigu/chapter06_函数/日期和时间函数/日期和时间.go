package main

import (
	"fmt"
	"strconv"
	"time"
)

// todo 日期和时间函数
func main() {

	//time.Time 类型 也是一种数据类型

	// 1、获取当前时间
	now := time.Now()
	fmt.Printf("now=%v now type=%T\n", now, now) //now=2023-03-06 13:55:11.35708 +0800 CST m=+0.001599601 now type=time.Time
	// 2、通过now可以获取到年月日、时分秒
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Printf("年=%v 月=%v 日=%v 时=%v 分=%v 秒=%v\n", year, int(month), day, hour, minute, second) //年=2023 月=March 日=6 时=15 分=35 秒=31

	// 3、格式化日期时间
	// 方式一：直接使用Printf和Sprintf,Sprintf可以将输出的字符串保存到一个变量中
	fmt.Printf("当前年月日: %d-%d-%d %d:%d:%d\n",
		now.Year(), now.Month(), now.Day(), now.Hour(),
		now.Minute(), now.Second())
	datestr := fmt.Sprintf("当前年月日: %d-%d-%d %d:%d:%d\n",
		now.Year(), now.Month(), now.Day(), now.Hour(),
		now.Minute(), now.Second())
	fmt.Println("datestr=", datestr)
	// 方式二：格式化的第二种方式：Format中只能写2006/01/02 15:04:05，也可以写其中一部分，只要数字不改变就行
	fmt.Printf(now.Format("2006/01/02 15:04:05\n")) //2023/03/06 15:44:46

	// 4、时间常量
	// 常量的作用：在程序中可用于获取指定时间单位的时间，比如向得到100毫秒：100*time.Millisecond
	/*const{
		Nanosecond Duration = 1 //纳秒
		Microsecond = 1000 * Nanosecond //微秒
		Millisecond = 1000 * Microsecond //毫秒
		Second = 1000 * Millisecond //秒
		Minute = 60 * Second //分
		Hour = 60 * Minute //小时
	}*/
	// 时间常量定义在time包中，使用time.Hour调用

	// 5、休眠
	fmt.Println("start...")
	time.Sleep(2000 * time.Millisecond) //休眠2000毫秒
	fmt.Println("end...")

	//每隔1秒打印一个数字
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}

	// 6、Unix() 秒 UnixNano() 纳秒 时间戳
	fmt.Println("unix时间戳", now.Unix())         //unix时间戳 1678089777
	fmt.Println("unixnano时间戳", now.UnixNano()) //unixnano时间戳 1678089777683501400

	// 7、案例：统计函数执行时间
	// 在执行test前，先获取当前的unix时间戳
	start := time.Now().Unix()
	test()
	end := time.Now().Unix()
	fmt.Printf("执行test所用时间:%v秒\n", end-start)
}
func test() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}
