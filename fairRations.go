/*
You are the benevolent ruler of Rankhacker Castle, and today you're
distributing bread. Your subjects are in a line, and some of them
already have some loaves. Times are hard and your castle's food stocks
are dwindling, so you must distribute as few loaves as possible
according to the following rules:
1. Every time you give a loaf of bread to some person i, you must also
give a loaf of bread to the person immediately in front of or behind them
in the line (i.e., persons i + 1 or i - 1).
2. After all the bread is distributed, each person must have an even
number of loaves.
Given the number of loaves already held by each citizen, find and print
the minimum number of loaves you must distribute to satisfy the two rules
above. If this is not possible, print NO.

Example
B = [4, 5, 6, 7]
We can first give a loaf to i = 3 and i = 4 so B = [4, 5, 7, 8].
Next we give a loaf to i = 2 and i = 3 and have B = [4, 6, 8, 8]
which satisfies our conditions.
All of the counts are now even numbers. We had to distribute 4 loaves.

Function Description
Complete the fairRations function in the editor below.
fairRations has the following parameter(s):
int B[N]: the numbers of loaves each persons starts with

Returns
string: the minimum number of loaves required, cast as a string, or 'NO'

Input Format
The first line contains an integer n, the number of subjects in the bread line.
The second line contains n space-separated integers B[i].

Constraints
2 <= n <= 1000
1 <= B[i] <= 10, where 0 <= i < n
*/

package main

import "fmt"

func fairRations(B []int32) string {
	// Write your code here
	var count int32 = 0

	for i := 0; i < len(B)-1; i++ {
		if B[i]%2 == 1 {
			count += 2
			B[i+1]++
		}
	}

	if B[len(B)-1]%2 == 1 {
		return "NO"
	}

	return fmt.Sprint(count)
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

	fmt.Println(fairRations(ar))
}
