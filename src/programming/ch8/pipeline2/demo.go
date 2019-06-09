//单向通道：chan <- int 只能发送的通道，<-chan int 只能接受的通道
package main

import "fmt"

func counter(out chan <- int)  {
	defer close(out)
	for x:= 0; x < 10; x++ {
		out <- x
	}
}

func squarer(in <-chan int, out chan <- int)  {
	defer close(out)
	for x := range in {
		out <- x*x
	}
}

func printer(in <-chan int)  {
	for x := range in {
		fmt.Println(x)
	}
}

func main()  {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(naturals, squares)
	printer(squares)
}

