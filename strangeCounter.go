/*
There is a strange counter. At the first second, it displays the number 3.
Each second, the number displayed by decrements by 1 until it reaches 1.
In next second, the timer resets to 2 * the initial number for the prior cycle
and continues counting down. The diagram below shows the counter values
for each time t in the first three cycles:

time value
1    3
2    2
3    1
4    6
5    5
6    4
7    3
8    2
9    1
10   12
11   11
12   10
...  ...
21   1
22   24
...  ...

Find and print the value displayed by the counter at time t.

Function Description
Complete the strangeCounter function in the editor below.
strangeCounter has the following parameter(s):
int t: an integer

Returns
int: the value displayed at time t.

Input Format
A single integer, the value of t.

Constraints
1 <= t <= 10^12

Subtask
1 <= t <= 10^5 for 60% of the maximum score.
*/

package main

import "fmt"

func strangeCounter(t int64) int64 {
	// Write your code here
	var counter int64 = 3

	for t > counter {
		t -= counter
		counter *= 2
	}

	return counter + 1 - t
}

func main() {
	var t int64

	fmt.Scanf("%v\n", &t)

	fmt.Println(strangeCounter(t))
}
