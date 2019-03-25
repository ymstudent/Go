package main

import (
	"fmt"
	"time"
)

//select语句使一个goroutine可以等待多个通信操作
//select 会阻塞到某个case可以继续执行为止，然后就会执行该case。如果同时有多个case都可以执行，那么就会随机选择一个执行

func fibonacci(c, quit chan int)  {
	x, y := 0, 1
	for  {
		select {
		case c <- x:
			x, y = y, x+y
		case <- quit:
			fmt.Println("quit")
			return
		default: //当 select 中的其它分支都没有准备好时，default 分支就会执行。为了在尝试发送或者接收时不发生阻塞，可使用 default 分支：
			fmt.Println("default")
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func main()  {
	c := make(chan int)
	quit := make(chan int)
	//开启一个goroutine读取c管道
	go func() {
		for i := 0 ; i < 10 ; i++ {
			fmt.Println(<-c)
		}
		//quit <- 0
	}()
	fibonacci(c, quit)
}
