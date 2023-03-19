/*
This is a simple challenge to get things started. Given a sorted
array arr and a number V, can you print the index location of V
in the array?

Example
arr = [1, 2, 3]
V = 3
Return 2 for a zero-based index array.

Function Description
Complete the introTutorial function in the editor below. It must
return an integer representing the zero-based index of V.
introTutorial has the following parameter(s):
int arr[n]: a sorted array of integers
int V: an integer to search for

Returns
int: the index of V in arr.

Input Format

The first line contains an integer, V, a value to search for.
The next line contains an integer, n, the size of arr.
The last line contains n space-separated integers,
each a value of arr[i] where 0 <= i < n.

Constraints
1 <= n <= 1000
-1000 <= V <= 1000, V E arr
V will occur in exactly once.
*/

package main

import "fmt"

func introTutorial(V int32, arr []int32) int32 {
	// Write your code here
	var index int = -1
	var left, right int = 0, len(arr)

	for left <= right && index == -1 {
		var middle = (left + right) / 2

		if V > arr[middle] {
			left = middle + 1
		} else if V < arr[middle] {
			right = middle - 1
		} else {
			index = middle
		}
	}

	return int32(index)
}

func IntroTutorial() {
	var n int
	var V, element int32
	var ar []int32 = []int32{}

	fmt.Scanf("%v\n", &V)

	fmt.Scanf("%v\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v\n", &element)
		ar = append(ar, element)
	}

	fmt.Println(introTutorial(V, ar))
}
