/*
The previous challenges covered Insertion Sort, which is a simple and
intuitive sorting algorithm with a running time of O(n^2). In these
next few challenges, we're covering a divide-and-conquer algorithm called
Quicksort (also known as Partition Sort). This challenge is a modified
version of the algorithm that only addresses partitioning.
It is implemented as follows:
Step 1: Divide
Choose some pivot element, p, and partition your unsorted array, arr,
into three smaller arrays: left, right, and equal, where each element in
left < p, each element in right > p, and each element in equal = p.

Example
arr = [5, 7, 4, 3, 8]
In this challenge, the pivot will always be at arr[0], so the pivot is 5.
arr is divided into left = {4, 3}, equal = {5}, and right = {7, 8}.
Putting them all together, you get {4, 3, 5, 7, 8}. There is a flexible
checker that allows the elements of left and right to be in any order.
For example, {3, 4, 5, 8, 7} is valid as well.
Given arr and p = arr[0], partition arr into left, right, and equal using
the Divide instructions above. Return a 1-dimensional array containing each
element in left first, followed by each element in equal, followed by each
element in right.

Function Description
Complete the quickSort function in the editor below.
quickSort has the following parameter(s):
int arr[n]: arr[0] is the pivot element

Returns
int[n]: an array of integers as described above

Input Format
The first line contains n, the size of arr.
The second line contains n space-separated integers arr[i] (the unsorted array).
The first integer, arr[0], is the pivot element, p.

Constraints
1 <= n <= 1000
-1000 <= arr[i] <= 1000 where 0 <= i < n
All elements are distinct.

Sample Input
STDIN       Function
-----       --------
5           arr[] size n =5
4 5 3 7 2   arr =[4, 5, 3, 7, 2]

Sample Output
3 2 4 5 7

Explanation
arr = [4, 5, 3, 7, 2] Pivot: p = arr[0] = 4.
left = {}; equal = {4}; right = {}
arr[1] = 5 > p, so it is added to right.
left = {}; equal = {4}; right = {5}
arr[2] = 3 < p, so it is added to left.
left = {3}; equal = {4}; right = {5}
arr[3] = 7 > p, so it is added to right.
left = {3}; equal = {4}; right = {5, 7}
arr[4] = 2 < p, so it is added to left.
left = {3, 2}; equal = {4}; right = {5, 7}
Return the array {3, 2, 4, 5, 7}.

The order of the elements to the left and right of 4 does not need
to match this answer. It is only required that 3 and 2 are to
the left of 4, and 5 and 7 are to the right.
*/

package main

import "fmt"

func quickSort(arr []int32) []int32 {
	// Write your code here
	var pivot int32 = arr[0]
	var left []int32 = []int32{}
	var equal []int32 = []int32{pivot}
	var right []int32 = []int32{}

	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		}

		if arr[i] == pivot {
			equal = append(equal, arr[i])
		}

		if arr[i] > pivot {
			right = append(right, arr[i])
		}
	}

	left = append(left, equal...)
	left = append(left, right...)

	return left
}

func main() {
	var n, element int32
	var ar []int32 = []int32{}

	fmt.Scanf("%v\n", &n)
	for i := 0; i < int(n); i++ {
		fmt.Scanf("%v\n", &element)
		ar = append(ar, element)
	}

	fmt.Println(quickSort(ar))
}
