package main

import "fmt"

func main() {
	fmt.Println(twoSum3([]int{3, 2, 4}, 6))
}

//返回他们给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并的数组下标。

/**
 * 自己写的
 * 耗时：36ms, 内存3.1mb
 */
func twoSum(nums []int, target int) []int {
	//1、取出一个初值。
	//2、遍历数组，如果初值+数组中的值为target返回
	//3、如果不是，则将初值移动到下一个，继续第二步
	length := len(nums)
	for i := 1; i < length; i++ {
		start := nums[i-1]
		s := nums[i:]
		fmt.Println(s)
		for k, v := range s {
			if (start + v) == target {
				return []int{i - 1, i + k}
			}
		}
	}
	return []int{}
}

/**
 * 先创建map，再遍历数组的操作
 * 耗时：8ms，内存：3.9mb
 */
func twoSum2(nums []int, target int) []int {
	//1、定义一个map,将数组中的k-v都存到map中。
	//2、遍历数组，target - 当前值，为剩下的数，直接去map中找就行，map查找元素的消耗为O(1)

	tmp := make(map[int]int)
	res := make([]int, 2)

	for k, v := range nums {
		tmp[v] = k
	}

	for k, v := range nums {
		res[0] = k
		value := target - v
		index, ok := tmp[value]
		if ok && index != k {
			res[1] = index
			break
		}
	}
	return res
}

/**
 * 边遍历数组，边创建map的操作
 * 耗时：8ms, 内存3.7。奇怪我看有人也是这么做的，耗时才4ms。这个因该比上面性能好啊，少了一次循环，估计是测试用例数组不够大吧
 */
func twoSum3(nums []int, target int) []int {
	tmp := make(map[int]int)

	for k, v := range nums {
		fmt.Println(tmp)
		value := target - v
		index, ok := tmp[value]
		fmt.Println(index)
		if ok {
			return []int{k, index}
		}
		tmp[v] = k
	}
	return []int{}
}
