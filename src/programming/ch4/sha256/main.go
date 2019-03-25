package main

import (
	"crypto/sha256"
	"fmt"
)

func main()  {
	//通过sha256加密散列算法计算出一个摘要，判断2条原始信息是否相同
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}
