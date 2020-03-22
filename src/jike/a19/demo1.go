package main

import (
	"errors"
	"fmt"
)


//error的常用方式：在函数的返回结果中申明一个error类型，在调用函数后，先判断err是否为nil
func echo(request string) (response string, err error) {
	if request == "" {
		err = errors.New("empty request")
		return
	}
	response = fmt.Sprintf("echo: %s", request)
	return
}

//判断一个错误值具体代表那一类错误的方式
//1、对于已知范围内的一系列错误值，一般使用类型断言表达式或类型switch语句来判断
//2、对于已有相应变量且类型相同的一系列错误值，一般直接使用判等操作来判断
//3、对于没有相应变量且类型未知的一系列错误值，只能使用其错误信息的字符串形式表示来做判断

func main()  {
	for _, req := range []string{"", "hello"} {
		fmt.Printf("request: %s\n", req)
		resp, err := echo(req)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			continue
		}
		fmt.Printf("response: %s\n", resp)
	}
}
