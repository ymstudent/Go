package main

import (
	"fmt"
)

func main()  {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
	fmt.Println(maxSubArray([]int{-2, -1, -1, -10}))
}


//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

/**
 * 思路：遍历数组，sum用来计算和，当sum小与等于0时，舍弃sum（将sum置为当前值），因为这部分sum无论存在于哪个子序列，
 * 都会让总和变小。用ans来记录最大和
 * 耗时：8ms，内存：3.4mb
 */
func maxSubArray(nums []int) int {
	sum := 0
	ans := nums[0]
	for _, v := range nums {
		if sum > 0 {
			sum += v
		} else {
			sum = v
		}
		if sum > ans {
			ans = sum
		}
	}

	return ans
}

/**
 * 分治法：建基于多项分支递归的一种很重要的算法范式。字面上的解释是“分而治之”，就是把一个复杂的问题分成两个或更多的相同或相似的子问题，
 * 直到最后子问题可以简单的直接求解，原问题的解即子问题的解的合并。（eg：快速排序，归并排序，傅立叶变换）
 */