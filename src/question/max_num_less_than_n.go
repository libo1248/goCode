package main

import (
	"math"
)

/*
请实现一个函数，输入一个正整数 n，返回与 n 组成数字完全相同，且小于 n 的最大整数。若不存在这样的数，返回 0。例如：
165 => 156
7123 => 3721
8918 => 8891
23 => 0
103 => 0 (前导0不作为组成数字，因此31不符合要求）
*/

func getMaxNum(n uint32) uint32 {
	// 解析出输入的每位数字
	var nums []uint32
	x := n
	for x > 0 {
		nums = append(nums, x%10)
		x = x / 10
	}

	// 个位数直接返回
	if len(nums) == 1 {
		return 0
	}

	// 是否每个低位都大于等于相邻的高位
	ok := true
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			ok = false
			break
		}
	}
	if ok {
		return 0
	}

	// 数组全排列
	numsSlice := permute(nums)

	var o, res uint32
	for _, ns := range numsSlice {
		o = 0

		// 最高位为0继续
		if ns[len(ns)-1] == 0 {
			continue
		}

		for i, u := range ns {
			o += u * uint32(math.Pow(10, float64(i)))
		}
		if o < n && o > res {
			res = o
		}
	}

	return res
}

var res [][]uint32

func permute(nums []uint32) [][]uint32 {
	res = [][]uint32{}
	backTrack(nums, len(nums), []uint32{})
	return res
}
func backTrack(nums []uint32, numsLen int, path []uint32) {
	//fmt.Println("nums:", nums)

	if len(nums) == 0 {
		p := make([]uint32, len(path))
		copy(p, path)
		res = append(res, p)
	}
	for i := 0; i < numsLen; i++ {
		cur := nums[i]
		path = append(path, cur)
		nums = append(nums[:i], nums[i+1:]...) //直接使用切片
		backTrack(nums, len(nums), path)
		nums = append(nums[:i], append([]uint32{cur}, nums[i:]...)...) //回溯的时候切片也要复原，元素位置不能变
		path = path[:len(path)-1]

	}
}
