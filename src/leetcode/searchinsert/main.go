package main

import "fmt"

func main()  {
	fmt.Println(searchInsert2([]int{1, 3, 5, 6}, 2))
}

//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//你可以假设数组中无重复元素。

/**
 * 耗时：4ms，内存：3.2mb
 * 一个无重复数字的排序数组，这就很简单了。直接遍历数组，如果v与target相等的值就直接返回k，如果没有，则找到比他大的v，在那一位的前面插入
 * 这样插入后，target也排在第k位，所以也返回k。
 */
func searchInsert(nums []int, target int) int {
	for k, v := range nums {
		if v >= target {
			return k
		}
	}
	return len(nums)
}

//错误的二分查找法
func searchInsert2(nums []int, target int) int {
	left , right := 0, len(nums) - 1
	for left < right {
		mid := (left + right) / 2 //注意：在极端情况下会出现整数溢出，最好的写法是通过位运算来解决：(left + right) >> 1
		if target == nums[mid] {
			return mid
		}
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return  left
}

//优化后的二分查找法
func searchInsert3(nums []int, target int) int {
	left , right := 0, len(nums) - 1
	for left < right {
		mid := (left + right) >> 1
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid //注意边界问题
		}
	}
	return  left
}
