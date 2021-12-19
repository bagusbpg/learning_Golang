/*
Given an array of bird sightings where every element represents a bird type id,
determine the id of the most frequently sighted type.
If more than 1 type has been spotted that maximum amount, return the smallest of their ids.

Example
ar = [1, 1, 2, 2, 3]
There are two each of types 1 and 2, and one sighting of type 3.
Pick the lower of the two types seen twice: type 1.

Function Description
Complete the migratoryBirds function in the editor below.
migratoryBirds has the following parameter(s):
int arr[n]: the types of birds sighted

Returns
int: the lowest type id of the most frequently sighted birds

Input Format
The first line contains an integer, n, the size of ar.
The second line describes ar as n space-separated integers,
each a type number of the bird sighted.

Constraints
5 <= n <= 2 * 10^5
It is guaranteed that each type is 1, 2, 3, 4, or 5.
*/

package main

import "fmt"

func migratoryBirds(arr []int32) int32 {
	// Write your code here
	var sightings map[int]int32 = map[int]int32{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
	}

	for _, value := range arr {
		switch value {
		case 1:
			sightings[1] += 1
		case 2:
			sightings[2] += 1
		case 3:
			sightings[3] += 1
		case 4:
			sightings[4] += 1
		default:
			sightings[5] += 1
		}
	}

	var max int32 = sightings[1]

	for i := 2; i <= 5; i++ {
		if sightings[i] > max {
			max = sightings[i]
		}
	}

	var lowest int32 = 5

	for i := range sightings {
		if sightings[i] == max {
			if int32(i) < lowest {
				lowest = int32(i)
			}
		}
	}

	return lowest
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

	fmt.Println(migratoryBirds(ar))
}
