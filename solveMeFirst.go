/*
Complete the function solveMeFirst to compute the sum of two integers.

Example
a = 7
b = 3
Return 10

Function description
Complete the solveMeFirst function in the editor below.
solveMeFirst has the following paramters:
int a: the first value
int b: the second value
Returns
int result: the sum of a and b

Constraints
1 <= a,b <= 1000
*/

package main

import "fmt"

func solveMeFirst(a uint32, b uint32) uint32 {
	return (a + b)
}

func main() {
	var a, b, result uint32
	fmt.Scanf("%v\n%v", &a, &b)
	result = solveMeFirst(a, b)
	fmt.Println(result)
}
