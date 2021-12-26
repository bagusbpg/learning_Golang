/*
Lily likes to play games with integers. She has created a new game
where she determines the difference between a number and its reverse.
For instance, given the number 12, its reverse is 21. Their difference is 9.
The number 120 reversed is 21, and their difference is 99.

She decides to apply her game to decision making. She will look at
a numbered range of days and will only go to a movie on a beautiful day.

Given a range of numbered days, [i...j] and a number k,
determine the number of days in the range that are beautiful.
Beautiful numbers are defined as numbers where |i - reverse(i)| is evenly divisible by k.
If a day's value is a beautiful number, it is a beautiful day.
Return the number of beautiful days in the range.

Function Description
Complete the beautifulDays function in the editor below.
beautifulDays has the following parameter(s):
int i: the starting day number
int j: the ending day number
int k: the divisor

Returns
int: the number of beautiful days in the range

Input Format
A single line of three space-separated integers describing
the respective values of i, j, and k.

Constraints
1 <= i <= j <= 2 * 10^6
1 <= k <= 2 * 10^9
*/

package main

import (
	"fmt"
	"strconv"
)

func beautifulDays(i int32, j int32, k int32) int32 {
	// Write your code here
	var count int32 = 0

	for day := i; day <= j; day++ {
		var str_day, str_day_reversed string = strconv.Itoa(int(day)), ""

		for n := len(str_day) - 1; n >= 0; n-- {
			str_day_reversed += string(str_day[n])
		}

		number, _ := strconv.Atoi(str_day_reversed)

		var int_day_reversed int32 = int32(number)

		if (int_day_reversed-day)%k == 0 {
			count++
		} else {
			continue
		}
	}

	return count
}

func main() {
	var start, end, k int32

	fmt.Scanf("%v\n", &start)
	fmt.Scanf("%v\n", &end)
	fmt.Scanf("%v\n", &k)

	fmt.Println(beautifulDays(start, end, k))
}
