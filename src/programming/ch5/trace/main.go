package main

import (
	"log"
	"time"
)

func main()  {
	bigSlowOperation()
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")() //如果不加圆括号，trace函数返回的方法不会执行
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}
