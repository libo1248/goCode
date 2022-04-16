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

func TestSortOrderNum(t *testing.T) {
	orders := []int{
		11221234,
		11211234,
		11121235,
		11221232,
		11121231,
	}
	sortOrders(orders)

	fmt.Println(orders)
}

func TestLrcCache(t *testing.T) {
	lru := NewLRUCache(4)
	lru.Put(1, 1)
	lru.Print()
	lru.Put(2, 2)
	lru.Print()
	lru.Put(3, 3)
	lru.Print()
	lru.Put(4, 4)
	lru.Print()
	lru.Put(1, 1)
	lru.Print()
	lru.Put(6, 6)
	lru.Print()
	lru.Get(6)
	lru.Print()
	lru.Put(2, 2)
	lru.Print()
	lru.Put(4, 4)
	lru.Print()
	lru.Put(1, 1)
	lru.Print()
	lru.Put(4, 4)
	lru.Print()
}
