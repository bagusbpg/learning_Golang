/*
An avid hiker keeps meticulous records of their hikes.
During the last hike that took exactly steps steps,
for every step it was noted if it was an uphill, U, or a downhill, D step.
Hikes always start and end at sea level, and each step up or down represents
a 1 unit change in altitude. We define the following terms:
1. A mountain is a sequence of consecutive steps above sea level,
starting with a step up from sea level and ending with a step down to sea level.
2. A valley is a sequence of consecutive steps below sea level,
starting with a step down from sea level and ending with a step up to sea level.
Given the sequence of up and down steps during a hike,
find and print the number of valleys walked through.

Example
steps = 8
path = [DDUUUUDD]
The hiker first enters a valley 2 units deep.
Then they climb out and up onto a mountain 2 units high.
Finally, the hiker returns to sea level and ends the hike.

Function Description
Complete the countingValleys function in the editor below.
countingValleys has the following parameter(s):
int steps: the number of steps on the hike
string path: a string describing the path

Returns
int: the number of valleys traversed

Input Format
The first line contains an integer steps, the number of steps in the hike.
The second line contains a single string path, of steps characters that describe the path.

Constraints
2 <= steps <= 10^6
path[i] E {UD}
*/

package main

import "fmt"

func countingValleys(steps int32, path string) int32 {
	// Write your code here
	var height, count int32 = 0, 0

	for i := 0; i < int(steps); i++ {
		if string(path[i]) == "D" {
			height--
		} else {
			height++
			if height == 0 {
				count++
			}
		}
	}

	return count
}

func CountingValleys() {
	var steps int32
	var path string

	fmt.Scanf("%v\n", &steps)
	fmt.Scanf("%v\n", &path)

	fmt.Println(countingValleys(steps, path))
}
