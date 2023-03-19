package main

import "fmt"

// Cho mảng và target là tổng của 2 số bất kỳ trong mảng, in ra vị trí 2 số bất kỳ đó
func twoSum(nums []int, target int) []int {
	var slice []int
	for i := range nums {
		for j := i + 1; j < len(nums)-1; j++ {
			if (nums[i] + nums[j]) == target {
				slice = append(slice, i, j)
				return slice
			}
		}
	}
	return nil
}
func bai1() {
	r := twoSum([]int{2, 7, 11, 15}, 13)
	fmt.Println(r)
}
