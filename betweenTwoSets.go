/*
There will be two arrays of integers. Determine all integers that satisfy
the following two conditions:
1. The elements of the first array are all factors of the integer being considered
2. The integer being considered is a factor of all elements of the second array
These numbers are referred to as being between the two arrays.
Determine how many such numbers exist.

Example
a = [2, 6]
b = [24, 36]
There are two numbers between the arrays: 6 and 12.
6%2 = 0, 6%6 = 0, 24%6 = 0 and 36%6 = 0 for the first value.
12%2 = 0, 12%6 = 0 and 24%12 = 0, 36%12 = 0 for the second value.
Return 2.

Function Description
Complete the getTotalX function in the editor below.
It should return the number of integers that are betwen the sets.
getTotalX has the following parameter(s):
int a[n]: an array of integers
int b[m]: an array of integers

Returns
int: the number of integers that are between the sets

Input Format
The first line contains two space-separated integers, n and m,
the number of elements in arrays a and b.
The second line contains n distinct space-separated integers a[i]
where 0 <= i < n.
The third line contains m distinct space-separated integers b[j]
where 0 <= j < m.

Constraints
1 <= n,m <= 10
1 <= a[i] <= 100
1 <= b[j] <= 100
*/

package main

import "fmt"

func getTotalX(a []int32, b []int32) int32 {
	// Write your code here
	var greatest_a, smallest_b, count int32 = 0, 100, 0

	for _, value := range a {
		if value > greatest_a {
			greatest_a = value
		}
	}

	for _, value := range b {
		if value < smallest_b {
			smallest_b = value
		}
	}

	var factor int32 = greatest_a
	var flag1, flag2 bool = true, true

	for factor <= smallest_b {
		for i := 0; i < len(a) && flag1; i++ {
			if factor%a[i] != 0 {
				flag1 = false
			}
		}

		if flag1 {
			for j := 0; j < len(b) && flag2; j++ {
				if b[j]%factor != 0 {
					flag2 = false
				}
			}

			if flag2 {
				count++
			} else {
				flag2 = true
			}
		} else {
			flag1 = true
		}

		factor += greatest_a
	}
	return count
}

func GetTotalX() {
	var n, m int
	var element int32
	var ar1, ar2 []int32 = []int32{}, []int32{}

	fmt.Scanf("%v\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v\n", &element)
		ar1 = append(ar1, element)
	}

	fmt.Scanf("%v\n", &m)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v\n", &element)
		ar2 = append(ar2, element)
	}

	fmt.Println(getTotalX(ar1, ar2))
}
