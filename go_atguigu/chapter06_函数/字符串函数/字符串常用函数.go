package main

import (
	"fmt"
	"strconv"
	"strings"
)

// todo 字符串常用的系统函数
func main() {
	// 1、len(str)：按字节统计字符串长度，有中文的话，一个汉字占3字节
	str := "hello威威"
	fmt.Println("str len=", len(str)) //11

	// 2、字符串遍历，同时处理有中文的问题：r := []rune(str)
	str2 := "hello北京"
	/*for i := 0; i < len(str2); i++ { //中文会出现乱码
		fmt.Printf("字符=%c\n", str2[i])
	}*/
	// 将str2转成切片,可以解决遍历中文乱码问题
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}

	// 3、字符串转整数：n,err := strconv.Atoi("12"),如果转不成功，会返回error
	n, err := strconv.Atoi("hell")
	if err != nil {
		fmt.Println("转换错误", err) //转换错误 strconv.Atoi: parsing "hell": invalid syntax
	} else {
		fmt.Println("转换成的结果是：", n)
	}

	// 4、整数转字符串：str = strconv.Itoa(222)
	str = strconv.Itoa(1234)
	fmt.Printf("str=%v type str=%T\n", str, str) //str=1234 type str=string

	// 5、字符串转 []byte : var bytes = []byte("hello go")
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes) //bytes=[104 101 108 108 111 32 103 111]

	// 6、[]byte 转 字符串
	str = string([]byte{97, 98, 99})
	fmt.Println("str=", str) //str= abc

	// 7、10进制转 2、8、16进制 str = strconv.FormatInt(123,2),2 => 8,16,返回对应的字符串
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制：%v\n", str) //123对应的二进制：1111011
	str = strconv.FormatInt(123, 16)
	fmt.Printf("123对应的十六进制：%v\n", str) //123对应的十六进制：7b

	// 8、查找子串是否在指定的字符串中：strings.Contains("seafood","foo")
	b := strings.Contains("seafood", "foo")
	fmt.Printf("b=%v\n", b) //b=true

	// 9、统计一个字符串有几个指定的子串：strings.Count("ceshi","e")
	num := strings.Count("ceshiee", "ie")
	fmt.Printf("num=%v\n", num) //num=1

	// 10、不区分大小写的字符串比较(==是区分字母大小写的)：strings.EqualFold("abc","ABC")
	b = strings.EqualFold("abc", "ABC") //b= true
	fmt.Println("b=", b)
	fmt.Println("abc" == "ABC") //false,==区分大小写

	// 11、返回子串在字符串中第一次出现的index值，如果没有则返回-1：
	// strings.Index("NLT_abc","abc")
	index := strings.Index("NLT_abcabc", "abc")
	fmt.Println("index=", index) //index= 4

	// 12、返回子串在字符串中最后一次出现的index值，如果没有则返回-1：
	// strings.LastIndex("NLT_abc","abc")
	index = strings.LastIndex("go golang", "go")
	fmt.Println("index=", index) //index= 3

	// 13、将指定的子串另外一个子串：strings.Replace("go golang","go","go语言",n) ,n可以指定你希望替换几个，如果n=-1表示替换全部
	str2 = "go golang" //替换完后，str2不会发生变化，只是替换后返回一个新的字符串
	str = strings.Replace(str2, "go", "java", 1)
	fmt.Println("str=", str)   //str= java golang
	fmt.Println("str2=", str2) //str= go golang

	// 14、按照指定的某个字符为分割标识，将一个字符串分成字符串数组：
	// strings.Split("hello,world,ok",",")
	strArr := strings.Split("hello,world,ok", ",")
	fmt.Println("strArr=", strArr) //strArr= [hello world ok]

	// 15、将字符串的字母进行大小写转换：
	//strings.ToLower("Go")
	//strings.ToUpper("Go")
	fmt.Println(strings.ToUpper("Hello,Go")) //HELLO,GO
	fmt.Println(strings.ToLower("Hello,Go")) //hello,go

	// 17、将字符串左右两边的空格去掉：strings.TrimSpace(" this ia an apple   ")
	str = strings.TrimSpace(" this ia an apple   ")
	fmt.Printf("str=%q\n", str) //str= this ia an apple

	// 18、将字符串左右两边指定的字符去掉
	// strings.Trim("! hello! "," !") //将左右两边 ! 和 " "去掉
	str = strings.Trim("! hello! ", " !")
	fmt.Printf("str=%q\n", str) //str="hello"

	// 19、将字符串左边指定的字符去掉
	// strings.Trim("! hello! "," !") //将左边 ! 和 " "去掉
	str = strings.TrimLeft("! hello! ", " !")
	fmt.Printf("str=%q\n", str) //str="hello! "

	// 20、将字符串右边指定的字符去掉
	// strings.Trim("! hello! "," !") //将右边 ! 和 " "去掉
	str = strings.TrimRight("! hello! ", " !")
	fmt.Printf("str=%q\n", str) //str="! hello"

	// 21、判断字符串是否以指定的字符串开头：
	// strings.HasPrefix("http://ww.NLT_abc.jpg", "http")
	b = strings.HasPrefix("http://ww.NLT_abc.jpg", "http")
	fmt.Println("b=", b) //b= true

	// 22、判断字符串是否以指定的字符串结尾：
	// strings.HasSuffix("NLT_abc.jpg", "jpg")
	b = strings.HasSuffix("NLT_abc.jpg", "jpg")
	fmt.Println("b=", b) //b= true
}
