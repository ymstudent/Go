//管道：一个的输出是另一个的输入
package main

import "fmt"

func main()  {
	naturals := make(chan int)
	squares := make(chan int)

	//counter
	go func() {
		defer close(naturals)
		for x := 0; x < 10; x++ {
			naturals <- x
		}
	}()

	//squares
	go func() {
		defer close(squares)
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x* x
		}
	}()

	//printer
	for x := range squares{
		fmt.Println(x)
	}
}
