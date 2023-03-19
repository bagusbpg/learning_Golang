/*
Given an array of integers, calculate the ratios of its elements that are positive, negative, and zero.
Print the decimal value of each fraction on a new line with 6 places after the decimal.

Note: This challenge introduces precision problems. The test cases are scaled to six decimal places,
though answers with absolute error of up to 10^-4 are acceptable.

Example
ar = [1, 1, 0, -1, -1]
There are n = 5 elements, two positive, two negative and one zero.
Their ratios are 2/5 = 0.400000, 2/5 = 0.400000 and 1/5 = 0.200000.
Results are printed as:
0.400000
0.400000
0.200000

Function Description
Complete the plusMinus function in the editor below.
plusMinus has the following parameter(s):
int arr[n]: an array of integers

Print
Print the ratios of positive, negative and zero values in the array.
Each value should be printed on a separate line with 6 digits after the decimal.
The function should not return a value.

Input Format

The first line contains an integer, n, the size of the array.
The second line contains n space-separated integers that describe ar[n].

Constraints
0 < n <= 100
-100 <= arr[i] <= 100

Output Format
Print the following 3 lines, each to 6 decimals:
1. proportion of positive values
2. proportion of negative values
3. proportion of zeros
*/

package main

import "fmt"

func plusMinus(arr []int32) {
	// Write your code here
	var plus, minus, zero, length float32 = 0, 0, 0, float32(len(arr))

	for _, value := range arr {
		if value > 0 {
			plus++
		} else if value < 0 {
			minus++
		} else {
			zero++
		}
	}

	fmt.Println(plus / length)
	fmt.Println(minus / length)
	fmt.Println(zero / length)
}

func PlusMinus() {
	var n int
	var element int32
	var ar []int32 = []int32{}

	fmt.Scanf("%v\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v\n", &element)
		ar = append(ar, element)
	}

	plusMinus(ar)
}
