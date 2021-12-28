/*
Given an array of long integers 'arr' and a number 'num', iterate through
the elements in arr and double the value of num whenever an element equals
num. Find the maximum possible value of num, knowing that arr can be
reordered before the iteration.

Example
arr = [1, 2, 4, 11, 12, 8]
num = 2
Iterating through arr:
arr num
    2
1   2
2   4
4   8
8   16
11  16
12  16
The maximal value of num = 16. Note that arr could have been reordered
before iterating.

Function Description
Complete the function doubleSize in the editor below.
doubleSize has the following parameter(s):
long int arr[n]: an array of long integers
long int num: the base long integer

Returns
long int: the maximal value of num

Constraints
1 <= n <= 10^6
0 <= arr[i] <= 10^16
0 <= num <= 10^4
*/

package main

import (
	"fmt"
	"sort"
)

func doubleSize(arr []int64, b int64) int64 {
	// Write your code here
	var int_arr []int = []int{}

	for _, value := range arr {
		int_arr = append(int_arr, int(value))
	}

	sort.Ints(int_arr)

	for _, value := range int_arr {
		if int64(value) == b {
			b *= 2
		}
	}

	return b
}

func main() {
	var n int
	var element, num int64
	var ar []int64 = []int64{}

	fmt.Scanf("%v\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v\n", &element)
		ar = append(ar, element)
	}

	fmt.Scanf("%v\n", &num)

	fmt.Println(doubleSize(ar, num))
}
