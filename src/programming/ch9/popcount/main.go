//使用sync.Once进行延迟初始化
package main

import (
	"fmt"
	"sync"
)

var loadTableOnce sync.Once
var pc [256]byte

func table()  {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func loadTable() [256]byte {
	loadTableOnce.Do(table)
	return pc
}

func popCount(x uint64) int {
	pc = loadTable()
	return int(pc[byte(x >> (0*8))] +
			pc[byte(x >> (1*8))] +
			pc[byte(x >> (1*8))] +
			pc[byte(x >> (2*8))] +
			pc[byte(x >> (3*8))] +
			pc[byte(x >> (4*8))] +
			pc[byte(x >> (5*8))] +
			pc[byte(x >> (6*8))] +
			pc[byte(x >> (7*8))])
}

func main()  {
	fmt.Println(popCount(11))
}
