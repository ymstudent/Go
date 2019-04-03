package main

import "fmt"

func main()  {
	arr := []int{9,4,5,3,7,1,6}
	Sort(arr)
	fmt.Println(arr)
}

func Sort(arr []int) {
	llen := len(arr)
	for i := 1; i < llen ; i++ {
		for j :=0 ; j < llen - i ; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
