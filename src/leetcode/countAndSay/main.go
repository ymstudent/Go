package main

import (
	"fmt"
	"strings"
)

func main()  {

}

//报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：
//1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
//1 被读作  "one 1"  ("一个一") , 即 11。
//11 被读作 "two 1s" ("两个一"）, 即 21。
//21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。
//给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。
//注意：整数顺序将表示为一个字符串。


/**
 * 思路：没有思路，题目我都看不懂，继上次对自己的数学能力产生怀疑后，我又对自己的语文能力产生了怀疑
 */
func countAndSay(n int) string {
	// 递归终止条件
	if n == 1 {
		return "1"
	}

	// 递归读取上一个字符串
	s := countAndSay(n - 1)

	// 结果数组，偶数位为字符个数，奇数位字符本身
	nums := []int{}

	// 前一个字符
	var prevByte byte

	// 遍历字符串
	for i := 0; i <= len(s) - 1; i++ {
		if prevByte != s[i] { // 当前字符和前一个字符不同，则往结果数组中，插入当前字符个数1和字符本身
			nums = append(nums, 1, int(s[i] - '0'))
			prevByte = s[i]
		} else { // 当前字符和前一个字符相同，则更新当前字符的个数
			nums[len(nums) - 2]++
		}
	}

	// 用字符串数组，比连接字符串更高效
	result := []string{}
	for i := 0; i <= len(nums) - 1; i++ {
		result = append(result, fmt.Sprintf("%d", nums[i]))
	}

	return strings.Join(result, "")
}