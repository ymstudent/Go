package main

import "fmt"

//发送者可以通过close关闭一个信道来表示没有需要发生的值了。接收者可以通过为接收表达式分配第二个参数来测试信道是否被关闭
//如果没有值可以接收且信道已经关闭，那么在执行完 v, 0k := <-ch之后，ok会被设置为false
//循环for i:= range c 会不断的从信道中接收值，直到它被关闭
//注意：只有发送者才能关闭信道，而接收者不能。向一个已经关闭的信道发送数据会触发一个panic
//另一个要注意的点是：信道与文件不同，通常情况下无需关闭它们。只有在必须告诉接收者不在有需要发送的值时才有必要关闭。例如终止一个range循环

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0 ; i < n ; i++  {
		c <- x
		x, y = y, x+y
	}

	close(c) //如果不关闭会报错：fatal error: all goroutines are asleep - deadlock! 原因：如果我们不关闭信道，range会一直
	           //从信道中接收数据，当缓冲区为空之后，接收方会阻塞，报出fatal error，所以这个错误实际上是主goroutine发出的
}

func main()  {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
