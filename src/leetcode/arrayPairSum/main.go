package main

import (
	"fmt"
)

func main()  {
	fmt.Println(quickSort([]int{1, 4, 3, 2}))
}


//给定长度为 2n 的数组, 你的任务是将这些数分成 n 对, 例如 (a1, b1), (a2, b2), ..., (an, bn) ，使得从1 到 n 的 min(ai, bi) 总和最大。

/**
 * 根据简单例子观察出来的逻辑：当按从小到达的顺序两两分组时，所有组最小数加起来的合最大。
 * 耗时：4228ms，内存：6.4mb
 * 按这个逻辑下来的情况不理想，冒泡的时间复杂度是O(n*logn)，下面的加法复杂度是O(n)。看来有更简单的逻辑在
 * 目前没有用到他给的关键字n，为什么要给一个长度为2n的数组?
 * 他妈的坑，看了一下其他人的思路和我一样，但选择的是快速排序，难道快排比冒泡快这么多？快排也是O(n*logn)啊
 */
func arrayPairSum(nums []int) int {
	length := len(nums)
	//先排序
	for i := 1; i < length; i ++ {
		for j := 0; j < length - i; j ++ {
			if nums[j] > nums[j + 1] {
				nums[j], nums[j + 1] = nums[j + 1], nums[j]
			}
		}
	}

	ret := 0
	for k, v := range nums {
		if (k % 2) == 0 {
			ret += v
		}
	}
	return ret
}

/**
 * 快速排序版本
 * 耗时：80ms，内存：6.4mb
 * 差距太大了
 */
func arrayPairSum2(nums []int) int {
	nums = quickSort(nums)
	ret := 0
	for k, v := range nums {
		if (k % 2) == 0 {
			ret += v
		}
	}
	return ret
}


/**
 * 快速排序步骤：
 * 1、挑选一个基准值，一般为数组第一位
 * 2、分割：重新排序数组，所有比基准值小的放在基准值前面，所有比基准值大的放在基准值后面
 * 3、递归排序子序列
 */
func quickSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	mid := data[0]
	head, tail := 0, len(data) - 1
	for i := 1; i <= tail; {
		//大于基准值的放到数组最后去
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		//小于或等于基准值的放到数组前面去
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	quickSort(data[:head])
	quickSort(data[head+1:])
	return data
}