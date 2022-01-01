/*
Given a sequence of integers a, a triplet (a[i], a[j], a[k]) is beautiful if:
i < j < k
a[j] - a[i] = a[k] - a[j] = d
Given an increasing sequence of integers and the value of d, count the number of
beautiful triplets in the sequence.

Example
arr = [2, 2, 3, 3, 4, 5]
d = 1
There are three beautiful triplets, by index: [i, j, k] = [0, 2, 3], [1, 2, 3],
[2, 3, 4]. To test the first triplet, arr[j] - arr[i] = 3 - 2 = 1 and arr[k] - arr[j] =
4 - 3 = 1.

Function Description
Complete the beautifulTriplets function in the editor below.
beautifulTriplets has the following parameters:
int d: the value to match
int arr[n]: the sequence, sorted ascending

Returns
int: the number of beautiful triplets

Input Format
The first line contains integer, n, the length of the sequence
The second line contains integer, d, the beautiful difference.
The third line contains n space-separated integers arr[i].

Constraints
1 <= n <= 10^4
1 <= d <= 20
0 <= arr[i] <= 2 * 10^4
arr[i] > arr[i - 1]
*/

package main

import "fmt"

func beautifulTriplets(d int32, arr []int32) int32 {
	// Write your code here
	var count int32 = 0

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr) && j <= i+2*int(d); j++ {
			if arr[j]-arr[i] == d {
				for k := j + 1; k < len(arr) && k <= j+int(d); k++ {
					if arr[k]-arr[j] == d {
						count++
					}
				}
			}
		}
	}

	return count
}

func main() {
	var n int
	var element, d int32
	var ar []int32 = []int32{}

	fmt.Scanf("%v\n", &n)
	fmt.Scanf("%v\n", &d)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v\n", &element)
		ar = append(ar, element)
	}

	fmt.Println(beautifulTriplets(d, ar))
}
