/*
Given an array of integers, find the longest subarray where
the absolute difference between any two elements is less than or equal to 1.

Example
a = [1, 1, 2, 2, 4, 4, 5, 5, 5]
There are two subarrays meeting the criterion: [1, 1, 2, 2] and [4, 4, 5, 5, 5].
The maximum length subarray has 5 elements.

Function Description
Complete the pickingNumbers function in the editor below.
pickingNumbers has the following parameter(s):
int a[n]: an array of integers

Returns
int: the length of the longest subarray that meets the criterion

Input Format

The first line contains a single integer n, the size of the array a.
The second line contains n space-separated integers, each an a[i].

Constraints
2 <= n <= 100
0 < a[i] < 100
The answer will be >= 2.
*/

package main

import (
	"fmt"
	"sort"
)

func pickingNumbers(a []int32) int32 {
	// Write your code here
	var mapNumbers map[int]int32 = map[int]int32{}
	var unique []int = []int{}
	var length, longest int32 = 0, 0

	for _, number := range a {
		if _, exist := mapNumbers[int(number)]; exist {
			mapNumbers[int(number)]++
		} else {
			mapNumbers[int(number)] = 1
			unique = append(unique, int(number))
		}
	}

	for _, count := range mapNumbers {
		if count > longest {
			longest = count
		}
	}

	sort.Ints(unique)

	for i := 0; i < len(unique)-1; i++ {
		if unique[i+1]-unique[i] <= 1 {
			length = mapNumbers[unique[i]] + mapNumbers[unique[i+1]]
			if length > longest {
				longest = length
			}
		}
	}

	return longest
}

func main() {
	var n int
	var element int32
	var ar []int32 = []int32{}

	fmt.Scanf("%v\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v\n", &element)
		ar = append(ar, element)
	}

	fmt.Println(pickingNumbers(ar))
}
