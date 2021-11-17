package main

import "fmt"

func main() {
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

func isValidSudoku(board [][]byte) bool {
	lineMap := make(map[int]map[byte]struct{})
	columnMap := make(map[int]map[byte]struct{})
	cellMap := make(map[int]map[byte]struct{})

	for i, bytes := range board {
		if _, ok := lineMap[i]; !ok {
			lineMap[i] = make(map[byte]struct{})
		}

		for j, b := range bytes {
			if _, ok := columnMap[j]; !ok {
				columnMap[j] = make(map[byte]struct{})
			}
			ij := (i/3)*3 + (j / 3)
			if _, ok := cellMap[ij]; !ok {
				cellMap[ij] = make(map[byte]struct{})
			}

			if b != '.' {
				// 行有重复
				if _, ok := lineMap[i][b]; ok {
					fmt.Println("line", i, j)
					return false
				}
				lineMap[i][b] = struct{}{}

				// 列有重复
				if _, ok := columnMap[j][b]; ok {
					fmt.Println("column", i, j)
					return false
				}
				columnMap[j][b] = struct{}{}

				// 小九宫块有重复
				if _, ok := cellMap[ij][b]; ok {
					fmt.Println("cell", i, j, ij)
					return false
				}
				cellMap[ij][b] = struct{}{}
			}
		}
	}

	return true
}
