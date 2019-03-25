//输出标准输入中出现次数大于1的行，前面是次数
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	//Scan函数在读到新行的时候返回true,在没有更多内容的时候返回false（linux中ctrl+D Windows中ctrl+z 回车）
	for input.Scan() {
		counts[input.Text()]++
	}
	//其实我觉得将下面这个for循环放入上面的while循环中更好理解
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
