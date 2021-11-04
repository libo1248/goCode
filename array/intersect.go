package main

import (
	"fmt"
	"sort"
)

/*
两个数组的交集
给定两个数组，编写一个函数来计算它们的交集。

示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2,2]
示例 2:

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[4,9]

说明：

输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
我们可以不考虑输出结果的顺序。
进阶：

如果给定的数组已经排好序呢？你将如何优化你的算法？
如果nums1的大小比nums2小很多，哪种方法更优？
如果nums2的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？

相关标签：数组 哈希表 双指针 二分查找 排序
*/

func main() {
	nums1, nums2 := []int{1, 1, 1, 2, 3, 3}, []int{1, 1, 2, 2, 3, 3}
	fmt.Println(intersect1(nums1, nums2))

	sort.Ints(nums1)
	sort.Ints(nums2)
	fmt.Println(intersect2(nums1, nums2))
}

func intersect1(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int)
	m2 := make(map[int]int)

	for _, num := range nums1 {
		m1[num]++
	}

	var res []int

	for _, num := range nums2 {
		count1 := m1[num]
		m2[num]++
		count2 := m2[num]
		if count2 <= count1 {
			res = append(res, num)
		}
	}

	return res
}

func intersect2(nums1 []int, nums2 []int) []int {
	var res []int

	if len(nums1) == 0 || len(nums2) == 0 {
		return res
	}

	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}

	return res
}
