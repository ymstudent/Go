package main

import "fmt"

//Buffered Channels
//channels是可以带换成的，将缓冲长度作为第二个参数提供给make来初始化一个带缓冲的channel：ch := make(chan int, 10)
//仅当channel的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接收方会阻塞

func main()  {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//ch <- 3 //此时缓冲区已经被填满，goroutine被阻塞，程序报错：fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch) //此时缓冲区为空，goroutine被阻塞，程序报错：fatal error: all goroutines are asleep - deadlock!
}
