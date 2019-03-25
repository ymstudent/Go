//输出命令行参数：优化版
package main

import (
	"fmt"
	"os"
)

func main()  {
	s, sep := "", ""

	for _, v := range os.Args {
		s += sep + v
		sep = " "
	}
	fmt.Println(s)
}
