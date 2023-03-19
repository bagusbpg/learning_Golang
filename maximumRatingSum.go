/*
Ryan is movie obsessed and has collected a list of movie quality ratings.
He wants to watch the largest contiguous list of movies with the highest
cumulative ratings possible. To do this, he must calculate the sum of all
contiguous subarrays in order to determine the maximum possible subarray sum.
For example, ratings are arr = [-1, 3, 4, -2, 5, -7]. We can see that the
highest value contiguous subarray runs from arr[1] - arr[4] and is 3 + 4 -2 + 5 = 10.

Function Description
Complete the function maximumSum in the editor below. It must return a long
integer denoting the maximum sum for any contiguous subarray in arr.
maximumSum has the following parameter(s):
arr[0...n-1]: an array of integers

Constraints
1 <= n <= 10^6
-10^7 <= arr[i] <= 10^7
*/

package main

import "fmt"

func maximumSum(arr []int32) int64 {
	// Write your code here
	var greatest int64 = 0
	var partialSum []int64 = []int64{int64(arr[0])}

	for i := 1; i < len(arr); i++ {
		var last int = len(partialSum) - 1

		if arr[i] >= 0 && arr[i-1] >= 0 {
			partialSum[last] += int64(arr[i])
		} else if arr[i] < 0 && arr[i-1] < 0 {
			partialSum[last] += int64(arr[i])
		} else {
			partialSum = append(partialSum, int64(arr[i]))
		}
	}

	if partialSum[0] < 0 {
		partialSum = partialSum[1:]
	}

	if len(partialSum)%2 == 1 {
		partialSum = append(partialSum, 0)
	}

	for i := 1; i <= len(partialSum)/2; i++ {
		for j := 0; j <= len(partialSum)-2*i; j += 2 {
			var sum int64

			for k := j; k <= j+2*(i-1); k++ {
				sum += partialSum[k]
			}

			if sum > greatest {
				greatest = sum
			}
		}
	}

	return greatest
}

func MaximumSum() {
	var n int
	var element int32
	var ar []int32 = []int32{}

	fmt.Scanf("%v\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v\n", &element)
		ar = append(ar, element)
	}

	fmt.Println(maximumSum(ar))
}
