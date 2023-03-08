package main

import "fmt"

//todo 结构体 结构体是一个值类型

//结构体声明
/*type 结构体名称 struct {
	field1 type //F大写的话就是公共的，其他包可以使用
	field2 type
}*/
//1、field如果是切片或者map，需要先make才能使用
//2、不同结构体变量的字段是独立的，互不影响，一个结构体变量字段更改，不影响另外一个

//创建结构体变量和访问结构体字段
//1、方式1：直接声明 var person Person
//2、方式2： var person Person = Person{}
//3、new new后生成的是一个指针 var person *Person = new(Person)
//4、{} var p *Person = &Person{}

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int              //指针
	slice  []int             //切片
	map1   map[string]string //map
}
type Cat struct {
	Name   string
	Age    int
	Color  string
	Scores [3]int
}

func main() {
	//创建一个Cat变量
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Color = "白色"
	cat1.Age = 2
	fmt.Println(cat1)
	fmt.Println("猫猫的信息如下：")
	fmt.Println("name=", cat1.Name)
	fmt.Println("age=", cat1.Age)
	fmt.Println("color=", cat1.Color)

	var person Person
	fmt.Println(person) //{ 0 [0 0 0 0 0] <nil> [] map[]}
	//结构体中字段是切片或map的话，要使用需要先make分配内存
	person.slice = make([]int, 5)
	fmt.Println(person) //{ 0 [0 0 0 0 0] <nil> [0 0 0 0 0] map[]}
	person.map1 = make(map[string]string, 3)
	fmt.Println(person) //{ 0 [0 0 0 0 0] <nil> [0 0 0 0 0] map[]}

	var p1 Person = Person{Name: "tom", Age: 12}
	fmt.Println(p1) //{tom 12 [0 0 0 0 0] <nil> [] map[]}
	var p2 *Person = new(Person)
	//因为p2是一个指针，因此标准的给字段赋值方式：
	(*p2).Name = "smith"
	(*p2).Age = 12
	fmt.Println(*p2)
	//上面的代码也可以这样写：p2.Name = "wyx"
	//原因：go的设计者为了程序员方便，底层会对p2.Name = "wyx" 进行处理，会给p2加上取值运算
	p2.Name = "wyx"
	p2.Age = 23
	fmt.Println(*p2)

	var p3 *Person = &Person{} //也可以在括号中直接赋值
	//因为p3是一个指针，因此标准的访问字段的方法：
	//(*p3).Name = "www"
	//go设计者为了方便程序员使用，也可以使用 p3.Name = "www"
	//和上面一样，底层会对p3.Name = "www"进行处理，加上(*p3)
	p3.Name = "weiyeuxin" //简化方式
	(*p3).Age = 34        //标准方式
	fmt.Println(p3)
}
