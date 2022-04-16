package main

import (
	"math/rand"
	"time"
)

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

func (this *Solution) Reset() []int {
	return this.nums
}

func (this *Solution) Shuffle() []int {
	rand.Seed(time.Now().Unix())
	res := make([]int, len(this.nums))
	copy(res, this.nums)

	m := len(res) - 1
	n := m
	for m >= 0 {
		i := int(rand.Int31n(int32(n)))
		res[m], res[i] = res[i], res[m]
		m--
	}

	return res
}
