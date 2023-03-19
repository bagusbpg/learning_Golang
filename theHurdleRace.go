/*
A video player plays a game in which the character competes in a hurdle race.
Hurdles are of varying heights, and the characters have a maximum height
they can jump. There is a magic potion they can take that will increase
their maximum jump height by 1 unit for each dose. How many doses of the potion
must the character take to be able to jump all of the hurdles.
If the character can already clear all of the hurdles, return 0.

Example
height = [1, 2, 3, 3, 2]
k = 1
The character can jump 1 unit high initially and must take 3 - 1 = 2
doses of potion to be able to jump all of the hurdles.

Function Description
Complete the hurdleRace function in the editor below.
hurdleRace has the following parameter(s):
int k: the height the character can jump naturally
int height[n]: the heights of each hurdle

Returns
int: the minimum number of doses required, always 0 or more

Input Format
The first line contains two space-separated integers n and k,
the number of hurdles and the maximum height the character can jump naturally.
The second line contains n space-separated integers height[i] where
0 <= i < n.

Constraints
1 <= n,k <= 100
1 <= height[i] <= 100
*/

package main

import "fmt"

func hurdleRace(k int32, height []int32) int32 {
	// Write your code here
	var greatest int32 = 0

	for _, value := range height {
		if value > greatest {
			greatest = value
		}
	}

	if greatest > k {
		return greatest - k
	} else {
		return 0
	}

}

func HurdleRace() {
	var n int
	var element, k int32
	var ar []int32 = []int32{}

	fmt.Scanf("%v\n", &n)
	fmt.Scanf("%v\n", &k)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v\n", &element)
		ar = append(ar, element)
	}

	fmt.Println(hurdleRace(k, ar))
}
