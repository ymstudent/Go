//输出其命令行参数
package main

import (
	"fmt"
	"os"
	"time"
)

func main()  {
	start := time.Now()
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("sepnd %.2fs\n", time.Since(start).Seconds())
}
