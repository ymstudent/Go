package main

import "fmt"

//由于goroutine在相同的内存中运行，因此在访问共享变量时必须进行同步。sync包提供了这种能力，不过这在GO中并不经常用到，我们通常使用channels(信道)来进行同步
//Channels（信道）是带有类型的管道，可以通过它用信道操作符<-来发生或者接受值（“箭头”就是数据流的方向）。
//ch <- v 将v发送至信道ch
//v := <-ch //从信道ch接收值并赋予变量v
//和切片、映射一样，信道在是使用前必须创建：ch := make(chan int)
//默认情况下，发送和接收操作在另一端准备好之前都会阻塞。这使得goroutine可以在没有显示的锁或竞态变量的情况下同步

//下面的例子中对切片中的数据进行求和，将任务分配给2个goroutine。一旦2个goroutine完成了他们的计算，它就能算出最终的结果

func sum(s []int, c chan int)  {
	sum := 0
	for _,v := range s {
		sum += v
	}
	c <- sum //将和送入c
}

func main()  {
	s := []int {7,2,8,-9,4,0}
	c := make(chan int) //创建一个int类型的管道
	go sum(s[:len(s)/2], c) //创建一个新的goroutine进行计算
	go sum(s[len(s)/2:], c) //创建一个新的goroutine进行计算
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
