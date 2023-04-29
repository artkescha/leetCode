package main

import "fmt"

func main() {
	// Input: nums = [1,3,5,6], target = 5
	nums := []int{1, 3, 5}
	fmt.Println(searchInsert(nums, 4))
}

func searchInsert(nums []int, target int) int {
	idx := recursiveSearchInsert(nums, target, len(nums), 1)
	if target != nums[idx] {
		min := nums[0]
		max := nums[len(nums)-1]
		if target < min {
			return 0
		}
		if target > max {
			return len(nums)
		}
	}
	return idx
}

func recursiveSearchInsert(nums []int, target, midIdx, coef int) int {
	if len(nums) == 0 {
		return midIdx
	}
	if len(nums) == 1 {
		return midIdx + coef*1
	}
	middle := len(nums) / 2
	if nums[middle] == target {
		return coef*midIdx + middle
	}
	if nums[middle] > target {
		return midIdx + coef*middle
	}
	if nums[middle] < target {
		return midIdx + coef*middle
	}

	if target < nums[middle] {
		return recursiveSearchInsert(nums[:middle], target, middle, -1)
	} else {
		return recursiveSearchInsert(nums[middle:], target, middle, 1)
	}
}
