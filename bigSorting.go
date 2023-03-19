/*
Consider an array of numeric strings where each string is a positive number
with anywhere from 1 to 10^6 digits. Sort the array's elements in non-decreasing,
or ascending order of their integer values and return the sorted array.

Example
unsorted = ['1', '200', '150', '3']
Return the array ['1', '3', '150', '200'].

Function Description
Complete the bigSorting function in the editor below.
bigSorting has the following parameter(s):
string unsorted[n]: an unsorted array of integers as strings

Returns
string[n]: the array sorted in numerical order

Input Format
The first line contains an integer, n, the number of strings in unsorted.
Each of the n subsequent lines contains an integer string, unsorted[i].

Constraints
1 <= n <= 2 * 10^5
Each string is guaranteed to represent a positive integer.
There will be no leading zeros.
The total number of digits across all strings in unsorted is between
1 and 10^6 (inclusive).
*/

package main

import (
	"fmt"
	"sort"
)

func bigSorting(unsorted []string) []string {
	// Write your code here
	var mapLength map[int][]string = map[int][]string{}

	for _, str := range unsorted {
		mapLength[len(str)] = append(mapLength[len(str)], str)
	}

	var digits []int = []int{}

	for key := range mapLength {
		digits = append(digits, key)
	}

	sort.Ints(digits)

	var result []string = []string{}

	for _, digit := range digits {
		sort.Strings(mapLength[digit])
		result = append(result, mapLength[digit]...)
	}

	return result
}

func BigSorting() {
	var n int
	var element string
	var ar []string = []string{}

	fmt.Scanf("%v\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v\n", &element)
		ar = append(ar, element)
	}

	fmt.Println(bigSorting(ar))
}
