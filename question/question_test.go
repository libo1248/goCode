package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMaxNumLessThenN(t *testing.T) {
	n := uint32(rand.Int())
	fmt.Println(n, getMaxNum(n))

}
