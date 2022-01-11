/*
Given a sorted list with an unsorted number e in the rightmost cell,
can you write some simple code to insert e into the array so that it
remains sorted?
Since this is a learning exercise, it won't be the most efficient way
of performing the insertion. It will instead demonstrate the brute-force
method in detail.
Assume you are given the array arr = [1, 2, 4, 5, 3] indexed 0...4.
Store the value of arr[4]. Now test lower index values successively
from 3 to 0 until you reach a value that is lower than arr[4], at arr[1]
in this case. Each time your test fails, copy the value at the lower index
to the current index and print your array. When the next lower indexed value
is smaller than arr[4], insert the stored value at the current index and
print the entire array.

Example
n = 5
arr = [1 2 4 5 3]
Start at the rightmost index. Store the value of arr[4] = 3. Compare this
to each element to the left until a smaller value is reached. Here are
the results as described:

1 2 4 5 5
1 2 4 4 5
1 2 3 4 5

Function Description
Complete the insertionSort1 function in the editor below.
insertionSort1 has the following parameter(s):
n: an integer, the size of arr
arr: an array of integers to sort

Returns
None: Print the interim and final arrays, each on a new line.
No return value is expected.

Input Format
The first line contains the integer n, the size of the array arr.
The next line contains n space-separated integers arr[0]...arr[n-1].

Constraints
1 <= n <= 1000
-10000 <= arr[i] <= 10000

Output Format
Print the array as a row of space-separated integers each time
there is a shift or insertion.
*/

package main

import "fmt"

func insertionSort1(n int32, arr []int32) {
	// Write your code here
	var flag bool = true
	var storage int32 = arr[n-1]

	for flag && n >= 2 {
		if arr[n-2] > storage {
			arr[n-1] = arr[n-2]
			for i := 0; i < len(arr); i++ {
				fmt.Printf("%v ", arr[i])
			}
			fmt.Println()
			n--
		} else {
			arr[n-1] = storage
			for i := 0; i < len(arr); i++ {
				fmt.Printf("%v ", arr[i])
			}
			fmt.Println()
			flag = false
		}
	}

	if flag {
		arr[0] = storage
		for i := 0; i < len(arr); i++ {
			fmt.Printf("%v ", arr[i])
		}
		fmt.Println()
	}
}

func main() {
	var n, element int32
	var ar []int32 = []int32{}

	fmt.Scanf("%v\n", &n)
	for i := 0; i < int(n); i++ {
		fmt.Scanf("%v\n", &element)
		ar = append(ar, element)
	}

	insertionSort1(n, ar)
}
