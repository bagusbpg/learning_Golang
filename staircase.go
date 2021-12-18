/*
Staircase detail
This is a staircase of size n = 4:

   #
  ##
 ###
####

Its base and height are both equal to n. It is drawn using # symbols and spaces.
The last line is not preceded by any spaces.
Write a program that prints a staircase of size n.

Function Description
Complete the staircase function in the editor below.
staircase has the following parameter(s):
int n: an integer

Print
Print a staircase as described above.

Input Format
A single integer, n, denoting the size of the staircase.

Constraints
0 < n <= 100.

Output Format
Print a staircase of size n using # symbols and spaces.
Note: The last line must have spaces in it.
*/

package main

import "fmt"

func staircase(n int32) {
	// Write your code here
	for i := 0; i < int(n); i++ {
		for j := 0; j < int(n)-i-1; j++ {
			fmt.Print(" ")
		}
		for l := int(n) - i - 1; l < int(n); l++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}

func main() {
	var n int32

	fmt.Scanf("%v\n", &n)
	staircase(n)
}
