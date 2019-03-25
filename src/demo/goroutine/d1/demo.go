package main

import (
	"fmt"
	"time"
)

//Go协程
//Goroutine是由Go运行时管理的轻量级线程：go f(x, y, z)会启动一个新的goroutine并执行：f(x, y, z)
//f,x,y和z的求值发生在当前的goroutine中，而f的执行发生在新的goroutine中

func say(s string)  {
	for i := 0 ; i < 5 ; i++ {
		time.Sleep(100 * time.Millisecond) //挂起100ms
		fmt.Println(s)
	}
}

func main()  {
	go say("World")
	say("Hello")
}
