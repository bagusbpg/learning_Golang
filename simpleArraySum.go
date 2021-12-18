/*
Given an array of integers, find the sum of its elements.

For example, if the array ar = {1, 2, 3}, 1 + 2 + 3 = 6, so return 6.

Function Description
Complete the simpleArraySum function in the editor below. It must return the sum of the array elements as an integer.

simpleArraySum has the following parameter(s):
ar: an array of integers

Input Format
The first line contains an integer, n, denoting the size of the array.
The second line contains n space-separated integers representing the array's elements.

Constraints
0 < n,ar[i] <= 1000

Output Format
Print the sum of the array's elements as a single integer.
*/

package main

import "fmt"

func simpleArraySum(ar []int32) int32 {
	// Write your code here
	var sum int32 = 0

	for _, element := range ar {
		sum += element
	}

	return sum
}

func main() {
	var n int
	var element int32
	var ar []int32

	fmt.Scanf("%v\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v\n", &element)
		ar = append(ar, element)
	}

	fmt.Println(simpleArraySum(ar))
}
