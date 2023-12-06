package main

import (
	"fmt"
)

func solveNQueens(n int) [][]string {
	res := [][]string{}
	curr := make([]int, 0, n)
	cols := make([]bool, n)
	diag1 := make([]bool, 2*n-1)
	diag2 := make([]bool, 2*n-1)
	placeQueen(&res, curr, 0, cols, diag1, diag2)
	return res
}

func placeQueen(
	res *[][]string,
	curr []int,
	row int,
	cols []bool,
	diag1 []bool,
	diag2 []bool,
) {
	n := len(cols)
	if row == n {
		*res = append(*res, getStringRows(curr))
		return
	}
	for col := 0; col < n; col++ {
		if !cols[col] && !diag1[col+row] && !diag2[row-col+n-1] {
			cols[col] = true
			diag1[col+row] = true
			diag2[row-col+n-1] = true
			curr = append(curr, col)
			placeQueen(res, curr, row+1, cols, diag1, diag2)
			curr = curr[:len(curr)-1]
			cols[col] = false
			diag1[col+row] = false
			diag2[row-col+n-1] = false
		}
	}
}

func getStringRows(cols []int) []string {
	res := make([]string, len(cols))
	for row := 0; row < len(cols); row++ {
		curr := make([]byte, len(cols))
		for col := 0; col < len(cols); col++ {
			if col == cols[row] {
				curr[col] = 'Q'
			} else {
				curr[col] = '.'
			}
		}
		res[row] = string(curr)
	}
	return res
}

/*  0 1 2 3
0 . . . .
1	. Q . .
2	. . . .
3	. . . .
*/

func main() {
	for i := 1; i <= 9; i++ {
		res := solveNQueens(i)
		fmt.Printf("n = %d, result:\n", i)
		for _, rows := range res {
			for _, row := range rows {
				fmt.Println(row)
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
