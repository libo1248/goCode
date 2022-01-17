package main

import (
	"fmt"
	"testing"
)

func TestMergeSlice(t *testing.T) {
	nums1 := []int{2, 0}
	m := 1
	nums2 := []int{1}
	n := 1
	mergeSlice(nums1, m, nums2, n)
	fmt.Println(nums1)
}
