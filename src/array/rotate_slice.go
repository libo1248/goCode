package main

/*
旋转数组
给你一个数组，将数组中的元素向右轮转 k个位置，其中k是非负数。

示例 1:
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]

示例2:
输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]


提示：
1 <= nums.length <= 105
-231 <= nums[i] <= 231 - 1
0 <= k <= 105


进阶：
尽可能想出更多的解决方案，至少有 三种 不同的方法可以解决这个问题。
你可以使用空间复杂度为O(1) 的原地算法解决这个问题吗？
相关标签：数组 数学 双指针
*/

func rotateSlice(nums []int, k int) {
	n := len(nums)
	if n == 0 || k <= 0 {
		return
	}

	k = k % n
	if k <= 0 {
		return
	}

	tmpNums := make([]int, n-k)
	copy(tmpNums, nums[:n-k])

	i := 0
	for _, num := range nums[n-k:] {
		nums[i] = num
		i++
	}
	for _, num := range tmpNums {
		nums[i] = num
		i++
	}
}
