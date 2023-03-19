/*
Given an array of integers, rearrange them so that the sum of
the absolute differences of all adjacent elements is minimized.
Then, compute the sum of those absolute differences.

Example
n = 5
ar = [1, 3, 3, 2, 4]
If the list is rearranged as ar' = [1, 2, 3, 3, 4], the absolute
differences are |1 - 2| = 1, |2 - 3| = 1, |3 - 3| = 0, |3 - 4| = 1.
The sum of those differences is 1 + 1 + 0 + 1 = 3.

Function Description
Complete the function minDiff in the editor below.
minDiff has the following parameter:
ar: an integer array

Returns
int: the sum of the absolute differences of adjacent
elements.

Constraints
2 <= n <= 10^5
0 <= ar[i] <= 10^9, where 0 <= i < n
*/

package main

import "fmt"

func minDiff(arr []int32) int32 {
	// Write your code here
	var smallest, greatest int32 = arr[0], arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
		} else if arr[i] > greatest {
			greatest = arr[i]
		}
	}

	return greatest - smallest
}

func MinDiff() {
	var n int
	var element int32
	var ar []int32 = []int32{}

	fmt.Scanf("%v\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v\n", &element)
		ar = append(ar, element)
	}

	fmt.Println(minDiff(ar))
}
