package main

import (
	"fmt"
	"testing"
)

func TestCountAndSay(t *testing.T) {
	for _, n := range []int{1, 2, 3, 4, 5} {
		fmt.Println(n, countAndSay(n))
	}
}

func TestFirstUniqChar(t *testing.T) {
	for _, s := range []string{"leetcode", "loveleetcode"} {
		fmt.Println(s, string(s[firstUniqChar(s)]))
	}
}

func TestIsAnagram(t *testing.T) {
	fmt.Println(isAnagram("ab", "baa"))
}

func TestNumReverse(t *testing.T) {
	for _, x := range []int{123, -654243123, 1, -1234567899} {
		fmt.Println(x, numReverse(x))
	}
}

func TestReverseString(t *testing.T) {
	for _, s := range [][]byte{
		[]byte("hello"),
		[]byte("test"),
	} {
		fmt.Printf("%s\n", s)
		reverseString(s)
		fmt.Printf("new: %s\n", s)
	}

}
