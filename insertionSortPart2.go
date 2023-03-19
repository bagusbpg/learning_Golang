/*
In Insertion Sort Part 1, you inserted one element into an array
at its correct sorted position. Using the same approach repeatedly,
can you sort an entire array?

Guideline: You already can place an element into a sorted array.
How can you use that code to build up a sorted array, one element
at a time? Note that in the first step, when you consider an array
with just the first element, it is already sorted since there's
nothing to compare it to.

In this challenge, print the array after each iteration of the
insertion sort, i.e., whenever the next element has been inserted
at its correct position. Since the array composed of just the first
element is already sorted, begin printing after placing the second
element.

Example.
n = 7
arr = [3, 4, 7, 5, 6, 2, 1]

Working from left to right, we get the following output:

3 4 7 5 6 2 1
3 4 7 5 6 2 1
3 4 5 7 6 2 1
3 4 5 6 7 2 1
2 3 4 5 6 7 1
1 2 3 4 5 6 7

Function Description
Complete the insertionSort2 function in the editor below.
insertionSort2 has the following parameter(s):
int n: the length of
int arr[n]: an array of integers

Prints
At each iteration, print the array as space-separated integers on its own line.

Input Format
The first line contains an integer, n, the size of arr.
The next line contains n space-separated integers arr[i].

Constraints
1 <= n <= 1000
-10000 <= arr[i] <= 10000, 0 <= i < n

Output Format
Print the entire array on a new line at every iteration.
*/

package main

import "fmt"

func insertionSort2(n int32, arr []int32) {
	// Write your code here
	for i := 1; i < int(n); i++ {
		var flag bool = true

		for j := 0; j < i && flag; j++ {
			if arr[i] < arr[j] {
				var newarr []int32 = []int32{}

				newarr = append(newarr, arr[:j]...)
				newarr = append(newarr, arr[i])
				newarr = append(newarr, arr[j:i]...)
				newarr = append(newarr, arr[i+1:]...)

				arr = newarr

				flag = false
			}
		}

		for j := 0; j < int(n); j++ {
			fmt.Printf("%v ", arr[j])
		}

		fmt.Println()
	}
}

func InsertionSort2() {
	var n int
	var element int32
	var ar []int32 = []int32{}

	fmt.Scanf("%v ", &n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%v ", &element)
		ar = append(ar, element)
	}

	insertionSort2(int32(n), ar)
}
