/*
You are given a square map as a matrix of integer strings.
Each cell of the map has a value denoting its depth. We will call
a cell of the map a cavity if and only if this cell is not
on the border of the map and each cell adjacent to it has
strictly smaller depth. Two cells are adjacent if they have
a common side, or edge.

Find all the cavities on the map and replace their depths with
the uppercase character X.

Example
grid = ['989', '191', '111']
The grid is rearranged for clarity:
989
191
111

Return:
989
1X1
111

The center cell was deeper than those on its edges: [8,1,1,1].
The deep cells in the top two corners do not share an edge with
the center cell, and none of the border cells is eligible.

Function Description
Complete the cavityMap function in the editor below.
cavityMap has the following parameter(s):
string grid[n]: each string represents a row of the grid

Returns
string{n}: the modified grid

Input Format
The first line contains an integer n, the number of rows and
columns in the grid. Each of the following n lines (rows) contains
n positive digits without spaces (columns) that represent the depth
at grid[row, column].

Constraints
1 <= n <= 100
*/

package main

import (
	"fmt"
	"strconv"
)

func cavityMap(grid []string) []string {
	// Write your code here
	var newgrid [][]int = [][]int{}
	var flags [][]bool = [][]bool{}

	for _, rowstr := range grid {
		var row []int = []int{}
		var flag []bool = []bool{}

		for _, str := range rowstr {
			digit, _ := strconv.Atoi(string(str))

			row = append(row, digit)
			flag = append(flag, false)
		}

		newgrid = append(newgrid, row)
		flags = append(flags, flag)
	}

	for i := 1; i < len(newgrid)-1; i++ {
		for j := 1; j < len(newgrid[i])-1; j++ {
			if newgrid[i-1][j] < newgrid[i][j] {
				if newgrid[i+1][j] < newgrid[i][j] {
					if newgrid[i][j-1] < newgrid[i][j] {
						if newgrid[i][j+1] < newgrid[i][j] {
							flags[i][j] = true
						}
					}
				}
			}
		}
	}

	var result []string

	for i := 0; i < len(newgrid); i++ {
		var str = ""

		for j := 0; j < len(newgrid[i]); j++ {
			if flags[i][j] {
				str += "X"
			} else {
				str += fmt.Sprint(newgrid[i][j])
			}
		}

		result = append(result, str)
	}

	return result
}

func main() {
	var n int
	var element string
	var ar []string = []string{}

	fmt.Scanf("%v\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v\n", &element)
		ar = append(ar, element)
	}

	fmt.Println(cavityMap(ar))
}
