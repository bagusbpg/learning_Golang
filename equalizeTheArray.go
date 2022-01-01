/*
Given an array of integers, determine the minimum number of elements
to delete to leave only elements of equal value.

Example
arr = [1, 2, 2, 3]
Delete the 2 elements 1 and 3 leaving arr = [2, 2]. If both twos
plus either the 1 or the 3 are deleted, it takes 3 deletions to leave
either [3] or [1]. The minimum number of deletions is 2.

Function Description
Complete the equalizeArray function in the editor below.
equalizeArray has the following parameter(s):
int arr[n]: an array of integers

Returns
int: the minimum number of deletions required

Input Format
The first line contains an integer n, the number of elements in arr.
The next line contains n space-separated integers arr[i].

Constraints
1 <= n,arr[i] <= 100
*/

package main

import "fmt"

func equalizeArray(arr []int32) int32 {
	// Write your code here
	var mapCount map[int32]int32 = map[int32]int32{}
	var number, most, deletions int32 = 0, 0, 0

	for _, value := range arr {
		if _, exist := mapCount[value]; exist {
			mapCount[value]++
		} else {
			mapCount[value] = 1
		}
	}

	for key, value := range mapCount {
		if value > most {
			most = value
			number = key
		}
	}

	for key, value := range mapCount {
		if key != number {
			deletions += value
		}
	}

	return deletions
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

	fmt.Println(equalizeArray(ar))
}
