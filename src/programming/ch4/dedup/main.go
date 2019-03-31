//使用map模拟集合功能
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	//使用map的键存储已经输入过的行，确保相同行不会重复输出
	seen := make(map[string]bool) //字符串
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedeup: %v\n", err)
		os.Exit(1)
	}
}
