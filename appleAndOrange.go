/*
Sam's house has an apple tree and an orange tree that yield an abundance of fruit.
Using the information given below, determine the number of apples and oranges that land on Sam's house.

In the diagram below:

    apple                           orange
     ***                             ***
    *****            _______        *****
     ***			/_______\        ***
      #            |         |        #
      #            |  house  |        #
----------------------------------------------
	  a            s         t        b

1. The house starts from s to t (both inclusive). The apple tree is to the left of the house,
and the orange tree is to its right.
2. Assume the trees are located on a single point, where the apple tree is at point a,
and the orange tree is at point b.
3. When a fruit falls from its tree, it lands d units of distance from its tree of origin
along the x-axis. A negative value of d means the fruit fell d units to the tree's left,
and a positive value of d means it falls d units to the tree's right. *

Given the value of d for m for apples and n oranges, determine how many apples and oranges
will fall on Sam's house (i.e., in the inclusive range [s, t])?

For example, Sam's house is between s = 7 and t = 10. The apple tree is located at a = 4
and the orange at b = 12. There are m = 3 apples and n = 3 oranges. Apples are thrown
apples = [2, 3, -4] units distance from a, and oranges are thrown oranges = [3, -2, -4]
units distance. Adding each apple distance to the position of the tree, they land at
[4+2, 4+3, 4-4] = [6, 7, 0]. Oranges land at [12+3, 12-2, 12-4] = [15, 10, 8].
One apple and two oranges land in the inclusive range 7 - 10 so we print
1
2

Function Description
Complete the countApplesAndOranges function in the editor below.
It should print the number of apples and oranges that land on Sam's house, each on a separate line.

countApplesAndOranges has the following parameter(s):
s: integer, starting point of Sam's house location.
t: integer, ending location of Sam's house location.
a: integer, location of the Apple tree.
b: integer, location of the Orange tree.
apples: integer array, distances at which each apple falls from the tree.
oranges: integer array, distances at which each orange falls from the tree.

Input Format
The first line contains two space-separated integers denoting the respective values of s and t.
The second line contains two space-separated integers denoting the respective values of a and b.
The third line contains two space-separated integers denoting the respective values of m and n.
The fourth line contains m space-separated integers denoting the respective distances that each apple
falls from point a.
The fifth line contains n space-separated integers denoting the respective distances that each orange
falls from point b.

Constraints
1 <= s,t,a,b,m,n <= 10^5
-10^5 <= d <= 10^5
a < s < t < b

Output Format
Print two integers on two different lines:
The first integer: the number of apples that fall on Sam's house.
The second integer: the number of oranges that fall on Sam's house.
*/

package main

import "fmt"

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	// Write your code here
	var count_a, count_b int32 = 0, 0

	for i := 0; i < len(apples); i++ {
		if apples[i]+a >= s && apples[i]+a <= t {
			count_a++
		}
	}

	for i := 0; i < len(oranges); i++ {
		if oranges[i]+b >= s && oranges[i]+b <= t {
			count_b++
		}
	}

	fmt.Println(count_a)
	fmt.Println(count_b)
}

func main() {
	var s, t, a, b, element int32
	var m, n int
	var apples, oranges []int32 = []int32{}, []int32{}

	fmt.Scanf("%v\n", &s)
	fmt.Scanf("%v\n", &t)
	fmt.Scanf("%v\n", &a)
	fmt.Scanf("%v\n", &b)
	fmt.Scanf("%v\n", &m)
	fmt.Scanf("%v\n", &n)

	for i := 0; i < m; i++ {
		fmt.Scanf("%v\n", &element)
		apples = append(apples, element)
	}

	for i := 0; i < n; i++ {
		fmt.Scanf("%v\n", &element)
		oranges = append(oranges, element)
	}

	countApplesAndOranges(s, t, a, b, apples, oranges)
}
