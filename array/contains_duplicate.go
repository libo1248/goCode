package main

import (
	"fmt"
	"sort"
)

/*
存在重复元素
给定一个整数数组，判断是否存在重复元素。

如果存在一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。

示例 1:
输入: [1,2,3,1]
输出: true

示例 2:
输入: [1,2,3,4]
输出: false

示例3:
输入: [1,1,1,3,3,4,3,2,4,2]
输出: true

相关标签：数组 哈希表 排序
*/

func main() {
	for _, nums := range [][]int{
		{1, 1, 1, 1, 2, 3, 3, 3, 3, 4, 5, 6, 6, 6, 6, 7, 7, 7, 7, 7},
		{},
		{1},
		{1, 2, 3, 4},
	} {
		fmt.Println(nums, containsDuplicate(nums))
		fmt.Println(nums, containsDuplicate2(nums))
	}
}

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}

	return false
}

func containsDuplicate2(nums []int) bool {
	m := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = struct{}{}
	}

	return false
}
