/*
A driver is driving on the freeway. The check engine light of his
vehicle is on, and the driver wants to get service immediately.
Luckily, a service lane runs parallel to the highway. It varies in
width along its length.
You will be given an array of widths at points along the road (indices),
then a list of the indices of entry and exit points. Considering each entry
and exit point pair, calculate the maximum size vehicle that can travel
that segment of the service lane safely.

Example
n = 4
width = [2, 3, 2, 1]
cases = [[1, 2], [2, 4]]
If the entry index, i = 1 and the exit, j = 2, there are two segment widths
of 2 and 3 respectively. The widest vehicle that can fit through both is 2.
If i = 2 and j = 4, the widths are [3, 2, 1] which limits vehicle width to 1.

Function Description
Complete the serviceLane function in the editor below.
serviceLane has the following parameter(s):
int n: the size of the width array
int cases[t][2]: each element contains the starting and ending indices for
a segment to consider, inclusive.

Returns
int[t]: the maximum width vehicle that can pass through each segment
of the service lane described.

Input Format
The first line of input contains two integers, n and t, where n denotes
the number of width measurements and t, the number of test cases.
The next line has n space-separated integers which represent the array width.
The next t lines contain two integers, i and j, where i is the start index
and j is the end index of the segment to check.

Constraints
2 <= n <= 100000
1 <= t <= 1000
0 <= i < j < n
2 <= j - i + 1 <= min(n, 1000)
1 <= width[k] <= 3, where 0 <= k < n
*/

package main

import "fmt"

func serviceLane(width []int32, cases [][]int32) []int32 {
	// Write your code here
	var result []int32 = []int32{}

	for _, lane := range cases {
		var smallest int32 = width[lane[0]]

		for i := lane[0] + 1; i <= lane[1]; i++ {
			if width[i] < smallest {
				smallest = width[i]
			}
		}

		result = append(result, smallest)
	}

	return result
}

func main() {
	var n, t int
	var element1, element2 int32
	var width []int32 = []int32{}
	var cases [][]int32 = [][]int32{}

	fmt.Scanf("%v\n", &n)
	fmt.Scanf("%v\n", &t)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v\n", &element1)
		width = append(width, element1)
	}
	for i := 0; i < t; i++ {
		fmt.Scanf("%v\n", &element1)
		fmt.Scanf("%v\n", &element2)
		cases = append(cases, []int32{element1, element2})
	}

	fmt.Println(serviceLane(width, cases))
}
