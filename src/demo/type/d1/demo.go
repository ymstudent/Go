package main

import "fmt"

//类型断言
//类型断言提供了访问接口值底层具体值的方式：
// t := i.(T)该语句断言接口i保存了具体类型T，并将其底层类型为T的值赋予变量t。若i并未保存T类型的值，该语句就会触发一个恐慌（panic）
//为了判断一个接口值是否保存了一个特定的类型，类型断言可返回2个值：其底层值以及一个报告断言是否成功的布尔值：t, ok := i.(T)
//如果i保存了一个T类型的值，则t为其底层值,ok=true，否则ok=false,t将为T类型的零值，程序不会触发恐慌。这种语法和检查映射中某个键是否存在具有相同之处

func main()  {
	var i interface {} //定义一个空接口
	i = "Hello" //保存一个类型为string的值
	//其实以上2步可以合并在一起：var i interface {} = "Hello"

	s := i.(string)
	fmt.Println(s) //Hello

	s, ok := i.(string)
	fmt.Println(s, ok) //Hello, true

	f, ok := i.(float64)
	fmt.Println(f, ok) //0， false

	f = i.(float64) //触发恐慌：panic: interface conversion: interface {} is string, not float64
}
