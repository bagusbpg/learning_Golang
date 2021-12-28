/*
In securities research, an analyst will look at a number of attributes for a stock.
One analyst would like to keep a record of the highest positive spread between
a closing price and the closing price on any prior day in history. Determine the
maximum positive spread for a stock given its price history. If the stock remains
flat or declines for the full period, return -1.

Example 0
px = [7, 1, 2, 5]
Calculate the positve difference between each price and its predecesors:
At the first quote, there is no earlier quote to compare to.
At the second quote, there was no earlier price that was lower.
At the third quote, the price is higher than the second quote: 2 - 1 = 1
For the fourth quote, the price is higher than the third and the second
quotes: 5 - 2 = 3, 5 - 1 = 4.
The maximum difference is 4.

Example 1
px = [7, 5, 3, 1]
The price declines each quote, so there is never a difference greater than 0.
In this case, return -1.

Function Description
Complete the function maxDifference in the editor below.
maxDifference has the following parameters:
int px[n]: an array of stock prices (quotes)

Returns
int: the maximum difference between two prices as described above

Constraints
1 <= n <= 10^5
-10^5 <= px[i] <= 10^5
*/

package main

import "fmt"

func maxDifference(px []int32) int32 {
	// Write your code here
	var points []int32 = []int32{px[0]}
	var greatest int32 = 0

	for i := 0; i < len(px)-3; i++ {
		if px[i] < px[i+1] && px[i+1] >= px[i+2] {
			points = append(points, px[i+1])
		} else if px[i] <= px[i+1] && px[i+1] > px[i+2] {
			points = append(points, px[i+1])
		} else if px[i] > px[i+1] && px[i+1] <= px[i+2] {
			points = append(points, px[i+1])
		} else if px[i] >= px[i+1] && px[i+1] < px[i+2] {
			points = append(points, px[i+1])
		}
	}

	points = append(points, px[len(px)-1])

	if len(points) == 2 {
		if px[len(px)-1] > px[0] {
			greatest = px[len(px)-1] - px[0]
		} else {
			greatest = -1
		}
	} else {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				if points[j]-points[i] > greatest {
					greatest = points[j] - points[i]
				}
			}
		}
	}

	return greatest
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

	fmt.Println(maxDifference(ar))
}
