/*
You are in charge of the cake for a child's birthday. You have decided the cake will have
one candle for each year of their total age. They will only be able to blow out
the tallest of the candles. Count how many candles are tallest.

Example
candles = [4, 4, 1, 3]
The maximum height candles are 4 units high. There are 2 of them, so return 2.

Function Description
Complete the function birthdayCakeCandles in the editor below.
birthdayCakeCandles has the following parameter(s):
int candles[n]: the candle heights

Returns
int: the number of candles that are tallest

Input Format
The first line contains a single integer, n, the size of candles[].
The second line contains n space-separated integers,
where each integer i describes the height of candles[i].

Constraints
1 <= n <= 10^5
1 <= candles[i] <= 10^7
*/

package main

import "fmt"

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	var tallest, count int32 = candles[0], 0

	// searching for the tallest
	for i := 1; i < len(candles); i++ {
		if candles[i] > tallest {
			tallest = candles[i]
		}
	}

	// counting the talesst candles
	for _, value := range candles {
		if value == tallest {
			count++
		}
	}

	return count
}

func BirthdayCakeCandles() {
	var n int
	var element int32
	var ar []int32 = []int32{}

	fmt.Scanf("%v\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v\n", &element)
		ar = append(ar, element)
	}

	fmt.Println(birthdayCakeCandles(ar))
}
