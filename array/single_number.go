package main

import (
	"fmt"
	"sort"
)

func singleNumber(nums []int) int {
	sort.Ints(nums)
	l := len(nums)
	if l == 1 {
		return nums[0]
	}

	if l == 2 {
		fmt.Println("此时没有出现一次的数字，或者2个数字都只出现了一次，不合题意")
	}

	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[i-1] && nums[i-1] != nums[i-2] {
			return nums[i-1] // 中间的数
		}

		if nums[i] == nums[i-1] && nums[i-1] != nums[i-2] && i == 2 {
			return nums[i-2] // 第一个数
		}
	}

	return nums[l-1] // 最后一个数
}
