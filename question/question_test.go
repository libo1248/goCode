package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMaxNumLessThenN(t *testing.T) {
	rand.Seed(time.Now().Unix())
	n := uint32(rand.Int())
	fmt.Println(n, getMaxNum(n))
}

func TestPermute(t *testing.T) {
	nums := []uint32{1, 2, 3}
	fmt.Println(nums)
	fmt.Println("---------------")
	fmt.Println(permute(nums))
}
