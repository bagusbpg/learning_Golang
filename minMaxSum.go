/*
Given five positive integers, find the minimum and maximum values that can be calculated by
summing exactly four of the five integers. Then print the respective minimum and maximum values
as a single line of two space-separated long integers.

Example
ar = [1, 3, 5, 7, 9]
The minimum sum is 1 + 3 + 5 + 7 = 16 and the maximum sum is 3 + 5 + 7 + 9 = 24.
The function prints
16 24

Function Description
Complete the miniMaxSum function in the editor below.
miniMaxSum has the following parameter(s):
ar: an array of 5 integers

Print
Print two space-separated integers on one line: the minimum sum and the maximum sum of 4
of 5 elements.

Input Format
A single line of five space-separated integers.

Constraints
1 <= ar[i] <= 10^9

Output Format
Print two space-separated long integers denoting the respective minimum and maximum values
that can be calculated by summing exactly four of the five integers.
(The output can be greater than a 32 bit integer.)
*/

package main

import "fmt"

func miniMaxSum(arr []int32) {
	// Write your code here
	var smallest, greatest, sum int64 = int64(arr[0]), int64(arr[0]), 0

	// searching the smallest and greatest value
	for i := 1; i < len(arr); i++ {
		if int64(arr[i]) < smallest {
			smallest = int64(arr[i])
		}
		if int64(arr[i]) > greatest {
			greatest = int64(arr[i])
		}
	}

	// calculating the sum
	for _, value := range arr {
		sum += int64(value)
	}

	fmt.Printf("%v %v\n", (sum - greatest), (sum - smallest))
}

func main() {
	var element int32
	var ar []int32 = []int32{}

	for i := 0; i < 5; i++ {
		fmt.Scanf("%v\n", &element)
		ar = append(ar, element)
	}

	miniMaxSum(ar)
}
