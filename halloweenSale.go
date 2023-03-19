/*
You wish to buy video games from the famous online video game store Mist.
Usually, all games are sold at the same price, p dollars. However, they are
planning to have the seasonal Halloween Sale next month in which you can
buy games at a cheaper price. Specifically, the first game will cost p dollars,
and every subsequent game will cost d dollars less than the previous one.
This continues until the cost becomes less than or equal to m dollars,
after which every game will cost m dollars. How many games can you buy during
the Halloween Sale?

Example
p = 20
d = 3
m = 6
s = 70
The following are the costs of the first 11, in order:
20, 17, 14, 11, 8, 6, 6, 6, 6, 6, 6
Start at P = 20 units cost, reduce that by d = 3 units each iteration until
reaching a minimum possible price, m = 6. Starting with s = 70 units of currency
in your Mist wallet, you can buy 5 games: 20 + 17 + 14 + 11 + 8 = 70.

Function Description
Complete the howManyGames function in the editor below.
howManyGames has the following parameters:
int p: the price of the first game
int d: the discount from the previous game price
int m: the minimum cost of a game
int s: the starting budget

Input Format
The first line of input contains integer p.
The second line of input contains integer d.
The third line of input contains integer m.
The last line of input contains integer s.

Constraints
1 <= m <= p <= 100
1 <= d <= 100
1 <= s <= 10^4
*/

package main

import "fmt"

func howManyGames(p int32, d int32, m int32, s int32) int32 {
	// Return the number of games you can buy
	var count int32 = 0

	for s >= p {
		s -= p

		if p-d >= m {
			p -= d
		} else {
			p = m
		}

		count++
	}

	return count
}

func HowManyGames() {
	var p, d, m, s int32

	fmt.Scanf("%v\n", &p)
	fmt.Scanf("%v\n", &d)
	fmt.Scanf("%v\n", &m)
	fmt.Scanf("%v\n", &s)

	fmt.Println(howManyGames(p, d, m, s))
}
