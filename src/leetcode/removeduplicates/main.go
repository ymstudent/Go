package main

import "fmt"

func main()  {
	fmt.Println(removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4}))
}

//给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
//不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成

/**
 * 思路：没有思路，不理解原地修改数组且在使用O(1)额外空间的条件下完成是啥意思？空间复杂度为O(1)?
 * 耗时：92ms，内存：7.9mb，时间复杂度：O(n)，空间复杂度：O(1)
 * 看得官方的解题思路，以时间换空间。
 * 设置2个指针i和j，i是快指针，j是慢指针。当nums[i] = nums[j]时，i++
 * 当nums[i] != nums[j]时，将i的值移动到j+1位，j++
 * 不过题目中说的是删除数组中的元素，所以被误导了一下，其实不是删除啊，只是统计了不同元素的个数，我说数组的删除操作为啥O（1）就可以了
 */
func removeDuplicates(nums []int) int {
	llen := len(nums)
	if llen < 2 {
		return 1
	}
	j := 0
	for i := 0; i < llen; i ++ {
		if nums[j] != nums[i] {
			nums[j + 1] = nums[i]
			j ++
		}
	}
	return j + 1
}