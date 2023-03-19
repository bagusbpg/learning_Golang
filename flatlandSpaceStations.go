/*
Flatland is a country with a number of cities, some of which have
space stations. Cities are numbered consecutively and each has a road of
1 km length connecting it to the next city. It is not a circular route,
so the first city doesn't connect with the last city. Determine
the maximum distance from any city to its nearest space station.

Example
n = 3
c = [1]
There are n = 3 cities and city 1 has a space station. They occur
consecutively along a route. City 0 is 1 - 0 = 1 unit away and
city 2 is 2 - 1 = 1 units away. City 1 is 0 units from its nearest
space station as one is located there. The maximum distance is 1.

Function Description
Complete the flatlandSpaceStations function in the editor below.
flatlandSpaceStations has the following parameter(s):
int n: the number of cities
int c[m]: the indices of cities with a space station

Returns
int: the maximum distance any city is from a space station

Input Format
The first line consists of two space-separated integers, n and m.
The second line contains m space-separated integers,
the indices of each city that has a space-station. These values
are unordered and distinct.

Constraints
1 <= n <= 10^5
1 <= m <= n
There will be at least 1 city with a space station.
No city has more than one space station.
*/

package main

import (
	"fmt"
	"sort"
)

func flatlandSpaceStations(n int32, c []int32) int32 {
	var newc []int = []int{}
	var distance int = 0

	for i := 0; i < len(c); i++ {
		newc = append(newc, int(c[i]))
	}

	sort.Ints(newc)

	if newc[0] != 0 {
		distance = newc[0] * 2
	}

	if newc[len(newc)-1] != int(n)-1 {
		if int(n)-1-newc[len(newc)-1] > distance/2 {
			distance = 2 * (int(n) - 1 - newc[len(newc)-1])
		}
	}

	for i := 1; i < len(newc); i++ {
		if newc[i]-newc[i-1] > distance {
			distance = newc[i] - newc[i-1]
		}
	}

	return int32(distance / 2)
}

func FlatlandSpaceStations() {
	var n, k, element int32
	var ar []int32

	fmt.Scanf("%v\n", &n)
	fmt.Scanf("%v\n", &k)
	for i := 0; i < int(k); i++ {
		fmt.Scanf("%v\n", &element)
		ar = append(ar, element)
	}

	fmt.Println(flatlandSpaceStations(n, ar))
}
