/*
There is a large pile of socks that must be paired by color.
Given an array of integers representing the color of each sock,
determine how many pairs of socks with matching colors there are.

Example
n = 7
ar = [1, 2, 1, 2, 1, 3, 2]
There is one pair of color 1 and one of color 2.
There are three odd socks left, one of each color.
The number of pairs is 2.

Function Description
Complete the sockMerchant function in the editor below.
sockMerchant has the following parameter(s):
int n: the number of socks in the pile
int ar[n]: the colors of each sock

Returns
int: the number of pairs

Input Format
The first line contains an integer n, the number of socks represented in ar.
The second line contains n space-separated integers, ar[i], the colors of the socks in the pile.

Constraints
1 <= n <= 100
1 <= ar[i] <= 100 where 0 <= i < n
*/

package main

import "fmt"

func sockMerchant(n int32, ar []int32) int32 {
	// Write your code here
	var count map[int32]int32 = map[int32]int32{}
	var pair int32 = 0

	count[ar[0]] = 1

	for i := 1; i < int(n); i++ {
		_, isPresent := count[ar[i]]

		if isPresent {
			count[ar[i]]++
		} else {
			count[ar[i]] = 1
		}
	}

	for _, value := range count {
		pair += (value / 2)
	}

	return pair
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

	fmt.Println(sockMerchant(int32(n), ar))
}
