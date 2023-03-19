/*
Given a sequence of nintegers, p(1), p(2), ..., p(n)
where each element is distinct and satisfies 1 <= p(x) <= n.
For each x where 1 <= x <= n, that is x increments from 1 to n,
find any integer y such that p(p(y)) = x and keep a history
of the values of y in a return array.

Example
p = [5, 2, 1, 3, 4]
Each value of x between 1 and 5, the length of the sequence, is analyzed as follows:
x = 1 -> p[3], p[4] = 3, so p[p[4]] = 1
x = 2 -> p[2], p[2] = 2, so p[p[2]] = 2
x = 3 -> p[4], p[5] = 4, so p[p[5]] = 3
x = 4 -> p[5], p[1] = 5, so p[p[1]] = 4
x = 5 -> p[1], p[3] = 1, so p[p[3]] = 5
The values for y are [4, 2, 5, 1, 3].

Function Description
Complete the permutationEquation function in the editor below.
permutationEquation has the following parameter(s):
int p[n]: an array of integers

Returns
int[n]: the values of y for all x in the arithmetic sequence 1 to n

Input Format
The first line contains an integer n, the number of elements in the sequence.
The second line contains n space-separated integers p[i] where 1 <= i <= n.

Constraints
1 <= n,p[i] <= 50, where 1 <= i <= n
Each element in the sequence is distinct.
*/

package main

import "fmt"

func permutationEquation(p []int32) []int32 {
	// Write your code here
	var mapNumber map[int32]int32 = map[int32]int32{}
	var index int32 = 1

	for _, value := range p {
		mapNumber[value] = index
		index++
	}

	var result []int32 = []int32{}

	for i := 1; i <= len(p); i++ {
		var y int32 = mapNumber[int32(i)]
		result = append(result, mapNumber[y])
	}

	return result
}

func PermutationEquation() {
	var n int
	var element int32
	var ar []int32 = []int32{}

	fmt.Scanf("%v\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v\n", &element)
		ar = append(ar, element)
	}

	fmt.Println(permutationEquation(ar))
}
