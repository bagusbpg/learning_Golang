/*
Happy Ladybugs is a board game having the following properties:

1. The board is represented by a string, b, of length n.
The ith character of the string, b[i], denotes the ith cell of
the board:
- If b[i] is an underscore (i.e., _), it means the ith cell of
the board is empty.
- If b[i] is an uppercase English alphabetic letter (ascii[A-Z]),
it means the ith cell contains a ladybug of color b[i].
- String b will not contain any other characters.

2. A ladybug is happy only when its left or right adjacent cell
(i.e., b[i +/- 1]) is occupied by another ladybug having the same
color.

3. In a single move, you can move a ladybug from its current
position to any empty cell.

Given the values of n and b, determine if it's possible to make all
the ladybugs happy. Return YES if all the ladybugs can be made happy
through some number of moves. Otherwise, return NO.

Example
b = [YYR_B_BR]
You can move the rightmost B and R to make b = [YYRRBB__] and all
the ladybugs are happy. Return YES.

Function Description
Complete the happyLadybugs function in the editor below.
happyLadybugs has the following parameters:
string b: the initial positions and colors of the ladybugs

Returns
string: either YES or NO

Input Format

The first line contains an integer n, the number of cells on
the board.
The second line contains a string b that describes the n cells
of the board.

Constraints
1 <= n <= 100
b[i] E {_, ascii[A-Z]}
*/

package main

import (
	"fmt"
	"strings"
)

func happyLadybugs(b string) string {
	// Write your code here
	if len(b) == 1 && b != "_" {
		return "NO"
	}

	if strings.Contains(b, "_") {
		var mapLadybugs map[string]int = map[string]int{}

		for _, value := range b {
			if string(value) != "_" {
				if _, exist := mapLadybugs[string(value)]; exist {
					mapLadybugs[string(value)]++
				} else {
					mapLadybugs[string(value)] = 1
				}
			}
		}

		for _, count := range mapLadybugs {
			if count == 1 {
				return "NO"
			}
		}
	} else {
		if b[0] != b[1] || b[len(b)-1] != b[len(b)-2] {
			return "NO"
		}

		for i := 1; i < len(b)-1; i++ {
			if b[i] != b[i-1] && b[i] != b[i+1] {
				return "NO"
			}
		}
	}
	return "YES"
}

func HappyLadyBugs() {
	var b string

	fmt.Scanf("%v\n", &b)

	fmt.Println(happyLadybugs(b))
}
