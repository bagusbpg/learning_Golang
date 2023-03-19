/*
Little Bobby loves chocolate. He frequently goes to his favorite 5&10 store,
Penny Auntie, to buy them. They are having a promotion at Penny Auntie.
If Bobby saves enough wrappers, he can turn them in for a free chocolate.

Example
n = 15
c = 3
m = 2
He has 15 to spend, bars cost 3, and he can turn in 2 wrappers to receive
another bar. Initially, he buys 5 bars and has 5 wrappers after eating them.
He turns in of 4 them, leaving him with 1, for 2 more bars. After eating those two,
he has 3 wrappers, turns in 2 leaving him with 1 wrapper and his new bar.
Once he eats that one, he has 2 wrappers and turns them in for another bar.
After eating that one, he only has 1 wrapper, and his feast ends. Overall,
he has eaten 5 + 2 + 1 + 1 = 9 bars.

Function Description
Complete the chocolateFeast function in the editor below.
chocolateFeast has the following parameter(s):
int n: Bobby's initial amount of money
int c: the cost of a chocolate bar
int m: the number of wrappers he can turn in for a free bar

Returns
int: the number of chocolates Bobby can eat after taking full advantage of
the promotion

Note: Little Bobby will always turn in his wrappers if he has enough to get
a free chocolate.

Input Format
The first line contains an integer n, the money to spend.
The second line contains an integer c, the cost of a chocolate.
The third line contains an integer m, the number of wrappers he can turn in
for a free chocolate.

Constraints
1 <= t <= 1000
2 <= n <= 10^5
1 <= c <= n
2 <= m <= n
*/

package main

import "fmt"

func chocolateFeast(n int32, c int32, m int32) int32 {
	// Write your code here
	var wrappers, bars int32 = n / c, n / c

	for wrappers >= m {
		bars += (wrappers / m)
		wrappers = wrappers/m + wrappers%m
	}

	return bars
}

func ChocolateFeast() {
	var n, c, m int32

	fmt.Scanf("%v\n", &n)
	fmt.Scanf("%v\n", &c)
	fmt.Scanf("%v\n", &m)

	fmt.Println(chocolateFeast(n, c, m))
}
