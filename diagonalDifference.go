/*
Given a square matrix, calculate the absolute difference between the sums of its diagonals.
For example, the square matrix arr is shown below:

1 2 3
4 5 6
9 8 9

The left-to-right diagonal = 1 + 5 + 9 = 15. The right to left diagonal = 3 + 5 + 9 = 17.
Their absolute difference is.

Function description
Complete the diagonalDifference function in the editor below.
diagonalDifference takes the following parameter:
int ar[n][m]: an array of integers

Return
int: the absolute diagonal difference

Input Format
The first line contains a single integer, n, the number of rows and columns in the square matrix ar.
Each of the next n lines describes a row, ar[i], and consists of n space-separated integers ar[i][j].

Constraints
-100 <= ar[i][j] <= 100
Output Format

Return the absolute difference between the sums of the matrix's two diagonals as a single integer.
*/

package main

import "fmt"

func diagonalDifference(ar [][]int32) int32 {
	// Write your code here
	var last int = len(ar) - 1
	var difference int32 = 0

	for i := 0; i <= last; i++ {
		difference += (ar[i][i] - ar[i][last-i])
	}

	if difference < 0 {
		difference = -difference
	}

	return difference
}

func main() {
	var n int
	var element int32
	var ar [][]int32 = [][]int32{{}, {}, {}}

	fmt.Scanf("%v\n", &n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scanf("%v\n", &element)
			ar[i] = append(ar[i], element)
		}
	}

	fmt.Println(diagonalDifference(ar))
}
