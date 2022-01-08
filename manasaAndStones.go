/*
Manasa is out on a hike with friends. She finds a trail of stones
with numbers on them. She starts following the trail and notices
that any two consecutive stones' numbers differ by one of two values.
Legend has it that there is a treasure trove at the end of the trail.
If Manasa can guess the value of the last stone, the treasure will be hers.

Example
n = 2
a = 2
b = 3
She finds 2 stones and their differences are a = 2 or b = 3. We know she
starts with a 0 stone not included in her count. The permutations of
differences for the two stones are [2, 2], [2, 3], [3, 2] or [3, 3].
Looking at each scenario, stones might have [2, 4], [2, 5], [3, 5] or [3, 6]
on them. The last stone might have any of 4, 5, or 6 on its face.

Compute all possible numbers that might occur on the last stone given
a starting stone with a 0 on it, a number of additional stones found,
and the possible differences between consecutive stones. Order the list
ascending.

Function Description
Complete the stones function in the editor below.
stones has the following parameter(s):
int n: the number of non-zero stones
int a: one possible integer difference
int b: another possible integer difference

Returns
int[]: all possible values of the last stone, sorted ascending

Input Format
The first line contains n, the number of non-zero stones found.
The second line contains a, one possible difference
The third line contains b, the other possible difference.

Constraints
1 <= n, a, b <= 10^3
*/

package main

import "fmt"

func stones(n int32, a int32, b int32) []int32 {
	// Write your code here
	if a == b {
		return []int32{a * (n - 1)}
	}

	if a > b {
		a, b = b, a
	}

	var stone int32 = a * (n - 1)
	var result []int32 = []int32{stone}

	var diff int32 = b - a

	for i := 1; i < int(n); i++ {
		stone += diff
		result = append(result, stone)
	}

	return result
}

func main() {
	var n, a, b int32

	fmt.Scanf("%v\n", &n)
	fmt.Scanf("%v\n", &a)
	fmt.Scanf("%v\n", &b)

	fmt.Println(stones(n, a, b))
}
