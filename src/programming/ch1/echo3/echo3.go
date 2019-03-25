package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main()  {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("sepnd %.2fs\n", time.Since(start).Seconds())
}
