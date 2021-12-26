/*
Watson likes to challenge Sherlock's math ability. He will provide
a starting and ending value that describe a range of integers,
inclusive of the endpoints.
Sherlock must determine the number of square integers within that range.
Note: A square integer is an integer which is the square of an integer, e.g.
1, 4, 9, 16, 25.

Example
a = 24
b = 49
There are three square integers in the range: 25, 36 and 49. Return 3.

Function Description
Complete the squares function in the editor below.
It should return an integer representing the number of square integers
in the inclusive range from a to b.
squares has the following parameter(s):
int a: the lower range boundary
int b: the upper range boundary

Returns
int: the number of square integers in the range

Input Format
The first line contains two space-separated integers, a and b,
the starting and ending integers in the ranges.

Constraints
1 <= a <= b <= 10^9
*/

package main

import (
	"fmt"
	"math"
)

func squares(a int32, b int32) int32 {
	// Write your code here
	var start, end int32 = int32(math.Sqrt(float64(a))), int32(math.Sqrt(float64(b)))

	if start*start < a {
		return end - start
	} else {
		return end - start + 1
	}
}

func main() {
	var a, b int32

	fmt.Scanf("%v\n", &a)
	fmt.Scanf("%v\n", &b)

	fmt.Println(squares(a, b))
}
