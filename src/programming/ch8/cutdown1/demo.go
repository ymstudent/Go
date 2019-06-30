//使用tick模拟火箭发射倒计时
package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Printf("Commencing countdownd.")
	tick := time.Tick(1 * time.Second) //tick返回一个通道。定期发送事件，每个事件的值为一个时间戳
	for countDown := 10; countDown > 0; countDown-- {
		fmt.Println(countDown)
		<- tick
	}
	launch()
}

func launch()  {
	fmt.Printf("launch !!!!")
}
