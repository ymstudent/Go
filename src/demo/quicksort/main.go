package main

import "fmt"

func main()  {
	fmt.Println(quickSort([]int{3,7,6,9,2,1}))
}

func quickSort(nums []int) []int {
	llen := len(nums)
	if llen <=1 {
		return nums
	}
	flag := nums[0] //设立标准值
	head, tail := 0, llen - 1
	for i := 1; i <= tail; {
		if nums[i] > flag { //大于标准值的放到后面
			nums[i], nums[tail] = nums[tail], nums[i]
			tail --
		} else { //小于或等于标准值的放到前面
			nums[i], nums[head] = nums[head], nums[i]
			head ++
			i ++
		}
		fmt.Println(i, head, tail, nums)
	}
	quickSort(nums[:head])
	quickSort(nums[head+1:])
	return nums
}

func maopao(nums []int) []int {
	llen := len(nums)
	//外层控制需要冒泡多少轮
	for i := 1; i < llen; i++ {
		//内层控制一个元素冒泡需要多少次
		for j := 0; j < llen - i; j++ {
			if nums[j] > nums[j+1] {
				nums[j] , nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
