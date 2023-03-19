/*
An integer d is a divisor of an integer n if the remainder of n / d is 0.
Given an integer, for each digit that makes up the integer determine
whether it is a divisor. Count the number of divisors occurring within the integer.

Example
n = 124
Check whether 1, 2 and 4 are divisors of 124. All 3 numbers divide evenly into 124
so return 3.
n = 111
Check whether 1, 1 and 1 are divisors of 111. All 3 numbers divide evenly into 111
so return 3.
n = 10
Check whether 1 and 0 are divisors of 10. 1 is, but 0 is not. Return 1.

Function Description
Complete the findDigits function in the editor below.
findDigits has the following parameter(s):
int n: the value to analyze

Returns
int: the number of digits in n that are divisors of n.

Input Format
The first line contain an integer, n.

Constraints
0 < n < 10^9
*/

package main

import "fmt"

func findDigits(n int32) int32 {
	// Write your code here
	var digit, number, count int32 = n % 10, n, 0

	for number != 0 {
		if digit != 0 && n%digit == 0 {
			count++
		}
		number /= 10
		digit = number % 10
	}

	return count
}

func FindDigits() {
	var n int32

	fmt.Scanf("%v\n", &n)

	fmt.Println(findDigits(n))
}
