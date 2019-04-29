package main

import (
	"flag"
	"fmt"
	"time"
)

//创建一个time.Duration类型的标志变量
var period = flag.Duration("period", 1*time.Second, "sleep period")

func main()  {
	flag.Parse()
	fmt.Printf("sleep for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
