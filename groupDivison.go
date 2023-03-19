/*
A university has admitted a group of n students with varying skill levels.
To better accomodate the students, the university has decided to create classes
tailored to the skill levels. A placement examination will return a skill level
that will be used to group the student, where levels[i] represents the skill
level of student i. All students within a group must have a skill level within
maxSpread, a specified range of one another. Determine the minimum number of
classes that must be formed.

Example:
n = 5
levels = [1, 4, 7, 3, 4]
maxSpread = 2

The students in any group must be within maxSpread = 2 levels of each other.
In this case, one optimal grouping is (1, 3), (4, 4), and (7). Another possible
grouping is (1), (3, 4, 4), (7). There is no way to form fewer than 3 groups.

Function Description
Complete the function groupDivison in the editor below.
groupDivison has the following parameter(s):
int levels[n]: the skill level for each student
int maxSpread: the maximum allowed skill level difference between any two members
of a group

Returns
int: the minimum number of groups that can be formed

Constraints
1 <= n <= 10^5
1 <= levels[i] <= 10^9 for every i (where 0 <= i <= n-1)
0 <= maxSpread <= 10^9
*/

package main

import (
	"fmt"
	"sort"
)

func groupDivision(levels []int32, maxSpread int32) int32 {
	// Write your code here
	var group []int = []int{}
	var newlevels []int = []int{}

	for _, value := range levels {
		newlevels = append(newlevels, int(value))
	}

	sort.Ints(newlevels)

	group = append(group, newlevels[0])

	for i := 1; i < len(newlevels); i++ {
		var last int = len(group) - 1

		if newlevels[i] > group[last]+int(maxSpread) {
			group = append(group, newlevels[i])
		}
	}

	return int32(len(group))
}

func GroupDivision() {
	var n int
	var element, maxSpread int32
	var ar []int32 = []int32{}

	fmt.Scanf("%v\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v\n", &element)
		ar = append(ar, element)
	}
	fmt.Scanf("%v\n", &maxSpread)

	fmt.Println(groupDivision(ar, maxSpread))
}
