/*
The distance between two array values is the number of indices between them.
Given a, find the minimum distance between any pair of equal elements in
the array. If no such value exists, return -1.

Example
a = [3, 2, 1, 2, 3]
There are two matching pairs of values: 3 and 2. The indices of the 3's are
i = 0 and j = 4, so their distance is d[i, j] = j - i = 4. The indices of the 2's
are i = 1 and j = 3, so their distance is d[i, j] = j - i = 2. The minimum
distance is 2.

Function Description
Complete the minimumDistances function in the editor below.
minimumDistances has the following parameter(s):
int a[n]: an array of integers

Returns
int: the minimum distance found or -1 if there are no matching elements

Input Format
The first line contains an integer n, the size of array a.
The second line contains n space-separated integers a[i].

Constraints
1 <= n <= 10^3
1 <= a[i] <= 10^5
*/

package main

import (
	"fmt"
)

func minimumDistances(a []int32) int32 {
	// Write your code here
	var distance int32 = int32(len(a))
	var mapIndex map[int32][]int32 = map[int32][]int32{}

	for index, value := range a {
		mapIndex[value] = append(mapIndex[value], int32(index))
	}

	for _, value := range mapIndex {
		if len(value) > 1 {
			for i := 1; i < len(value); i++ {
				if value[i]-value[i-1] < distance {
					distance = value[i] - value[i-1]
				}
			}
		}
	}

	if distance != int32(len(a)) {
		return distance
	} else {
		return -1
	}
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

	fmt.Println(minimumDistances(ar))
}
