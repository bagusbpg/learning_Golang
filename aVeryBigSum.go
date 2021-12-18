/*
In this challenge, you are required to calculate and print the sum of the elements in an array, keeping in mind that some of those integers may be quite large.

Function Description
Complete the aVeryBigSum function in the editor below. It must return the sum of all array elements.
aVeryBigSum has the following parameter(s):
int ar[n]: an array of integers .

Return
long: the sum of all array elements

Input Format
The first line of the input consists of an integer n.
The next line contains n space-separated integers contained in the array.

Output Format
Return the integer sum of the elements in the array.

Constraints
1 <= n <= 10
0 <= ar[i] <= 10^10
*/

package main

import "fmt"

func aVeryBigSum(ar []int64) int64 {
	//Write your code here
	var sum int64 = 0

	for _, element := range ar {
		sum += element
	}

	return sum
}

func main() {
	var n int
	var element int64
	var ar []int64 = []int64{}

	fmt.Scanf("%v\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v\n", &element)
		ar = append(ar, element)
	}

	fmt.Println(aVeryBigSum(ar))
}
