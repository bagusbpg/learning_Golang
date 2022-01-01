/*
Taum is planning to celebrate the birthday of his friend, Diksha.
There are two types of gifts that Diksha wants from Taum: one is black
and the other is white. To make her happy, Taum has to buy b black gifts and
w white gifts. The cost of each black gift is bc units. The cost of
every white gift is wc units. The cost to convert a black gift into white gift
or vice versa is z units.
Determine the minimum cost of Diksha's gifts.

Example
b = 3
w = 5
bc = 3
wc = 4
z = 1
He can buy a black gift for 3 and convert it to a white gift for 1,
making the total cost of each white gift 4. That matches the cost of a white gift,
so he can do that or just buy black gifts and white gifts. Either way,
the overall cost is 3 * 3 + 5 * 4 = 29.

Function Description
Complete the function taumBday in the editor below.
It should return the minimal cost of obtaining the desired gifts.
taumBday has the following parameter(s):
int b: the number of black gifts
int w: the number of white gifts
int bc: the cost of a black gift
int wc: the cost of a white gift
int z: the cost to convert one color gift to the other color

Returns
int: the minimum cost to purchase the gifts

Input Format
The first line contains the values of integer b.
The second linfe contains the values of integer w.
The first line contains the values of integer bc.
The second linfe contains the values of integer wc.
The second linfe contains the values of integer z.

Constraints
0 <= b, w, bc, wc, z <= 10^9
*/

package main

import "fmt"

func taumBday(b int32, w int32, bc int32, wc int32, z int32) int64 {
	// Write your code here
	var cost int64

	if bc > wc+z {
		cost = int64(b)*(int64(wc)+int64(z)) + int64(w)*int64(wc)
	} else if wc > bc+z {
		cost = int64(b)*int64(bc) + int64(w)*(int64(bc)+int64(z))
	} else {
		cost = int64(b)*int64(bc) + int64(w)*int64(wc)
	}

	return cost
}

func main() {
	var b, w, bc, wc, z int32

	fmt.Scanf("%v\n", &b)
	fmt.Scanf("%v\n", &w)
	fmt.Scanf("%v\n", &bc)
	fmt.Scanf("%v\n", &wc)
	fmt.Scanf("%v\n", &z)

	fmt.Println(taumBday(b, w, bc, wc, z))
}
