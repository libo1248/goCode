package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	for _, nums := range [][]int{
		{1, 1, 1, 1, 2, 3, 3, 3, 3, 4, 5, 6, 6, 6, 6, 7, 7, 7, 7, 7},
		{},
		{1},
		{1, 2, 3, 4},
	} {
		fmt.Println(nums, containsDuplicate(nums))
		fmt.Println(nums, containsDuplicate2(nums))
	}
}

func TestFirstMissingPositive(t *testing.T) {
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
}

func TestIntersect(t *testing.T) {
	nums1, nums2 := []int{1, 1, 1, 2, 3, 3}, []int{1, 1, 2, 2, 3, 3}
	fmt.Println(intersect1(nums1, nums2))

	sort.Ints(nums1)
	sort.Ints(nums2)
	fmt.Println(intersect2(nums1, nums2))
}

func TestIsValidSudoku(t *testing.T) {
	for _, sudoku := range [][][]byte{
		{
			{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '2'},
		},
		{
			{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		},
	} {
		fmt.Println(sudoku, isValidSudoku(sudoku))
	}
}

func TestMaxProfit(t *testing.T) {
	for _, prices := range [][]int{
		{1, 2, 3, 4, 5, 7},
		{},
		{1},
	} {
		fmt.Println(maxProfit(prices))
	}
}

func TestMoveZeroes(t *testing.T) {
	nums := []int{1, 2, 0, 0, 5, 4, 3, 0, 9, 0}
	moveZeroes(nums)
	fmt.Println(nums)
}

func TestPlusOne(t *testing.T) {
	fmt.Println(plusOne([]int{8, 9, 9, 9, 9, 9}))
}

func TestRemoveDuplicates(t *testing.T) {
	for _, nums := range [][]int{
		{1, 1, 1, 1, 2, 3, 3, 3, 3, 4, 5, 6, 6, 6, 6, 7, 7, 7, 7, 7},
		{},
		{1},
		{1, 2, 3, 4},
	} {
		n := removeDuplicates(nums)
		fmt.Println(nums, nums[:n])
	}

}

func TestRotateMatrix(t *testing.T) {
	for _, matrix := range [][][]int{
		{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
		{{1}},
		{{1, 2}, {3, 4}},
	} {
		fmt.Println(matrix)
		rotateMatrix(matrix)
		fmt.Println("new:", matrix)
	}
}

func TestRotate(t *testing.T) {
	for _, nums := range [][]int{
		{},
		{1, 2},
		{1, 2, 3, 4, 5, 6, 7},
	} {
		n := 1
		fmt.Println("nums", nums, n)
		rotateSlice(nums, n)
		fmt.Println("newNums", nums)
	}

}

func TestConstructor(t *testing.T) {
	obj := Constructor([]int{1, 2})
	fmt.Println(obj.Reset())
	fmt.Println(obj.Shuffle())
}

func TestTwoSum(t *testing.T) {
	nums1, target1 := []int{2, 7, 11, 15}, 9
	fmt.Println(nums1, target1, twoSum(nums1, target1))
	nums2, target2 := []int{3, 2, 4}, 6
	fmt.Println(nums2, target2, twoSum(nums2, target2))
	nums3, target3 := []int{3, 3}, 6
	fmt.Println(nums3, target3, twoSum(nums3, target3))
}

func TestSingleNumber(t *testing.T) {
	nums := []int{1, 1, 2, 2, 7, 7, 11}
	fmt.Println(nums, singleNumber(nums))
}
