package main

import "fmt"

/*
移动零
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。

相关标签：数组 双指针
*/

func main() {
	nums := []int{1, 2, 0, 0, 5, 4, 3, 0, 9, 0}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
	}

	for i < len(nums) {
		nums[i] = 0
		i++
	}
}
