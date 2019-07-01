//select多路复用
package main

import (
	"fmt"
	"os"
	"time"
)

func main()  {
	fmt.Printf("Commencing countdown, Press return to abort.")
	tick := time.Tick(1 * time.Second)
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	for countDown := 10; countDown > 0; countDown-- {
		fmt.Println(countDown)
		select {
		case <-tick:
			//不执行任何操作
		case <-abort:
			fmt.Println("Launch aborted")
			return
		}
	}
	launch()
}

func launch()  {
	fmt.Printf("launch !!!!")
}
