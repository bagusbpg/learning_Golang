/*
You are given a number of sticks of varying lengths. You will iteratively
cut the sticks into smaller sticks, discarding the shortest pieces
until there are none left. At each iteration you will determine the length of
the shortest stick remaining, cut that length from each of the longer sticks
and then discard all the pieces of that shortest length.
When all the remaining sticks are the same length, they cannot be shortened,
so discard them.

Given the lengths of n sticks, print the number of sticks that are left
before each iteration until there are none left.

Example
arr = [1, 2, 3]
The shortest stick length is 1, so cut that length from the longer two and
discard the pieces of length 1. Now the lengths are arr = [1, 2]. Again,
the shortest stick is of length 1, so cut that amount from the longer stick
and discard those pieces. There is only one stick left, arr = [1], so discard
that stick. The number of sticks at each iteration are answer = [3, 2, 1].

Function Description
Complete the cutTheSticks function in the editor below. It should return
an array of integers representing the number of sticks before each cut operation
is performed.
cutTheSticks has the following parameter(s):
int arr[n]: the lengths of each stick

Returns
int[]: the number of sticks after each iteration

Input Format
The first line contains a single integer n, the size of .
The next line contains n space-separated integers, each an arr[i],
where each value represents the length of the ith stick.

Constraints
1 <= n <= 1000
1 <= arr[i] <= 1000
*/

package main

import (
	"fmt"
	"sort"
)

func cutTheSticks(arr []int32) []int32 {
	// Write your code here
	var newarr []int = []int{}
	var result []int32 = []int32{int32(len(arr))}

	for _, value := range arr {
		newarr = append(newarr, int(value))
	}

	sort.Ints(newarr)

	for i := 1; i < len(newarr); i++ {
		if newarr[i] > newarr[i-1] {
			result = append(result, int32(len(arr)-i))
		}
	}

	return result
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

	fmt.Println(cutTheSticks(ar))
}
