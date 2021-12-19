/*
A person wants to determine the most expensive computer keyboard and USB drive
that can be purchased with a give budget. Given price lists for keyboards
and USB drives and a budget, find the cost to buy them.
If it is not possible to buy both items, return -1.

Example
b = 60
keyboards = [40, 50, 60]
drives = [5, 8, 12]
The person can buy a 40 keyboard + 12 USB drive = 52, or a 50 keyboard + 8 USB drive = 58.
Choose the latter as the more expensive option and return 58.

Function Description
Complete the getMoneySpent function in the editor below.
getMoneySpent has the following parameter(s):
int keyboards[n]: the keyboard prices
int drives[m]: the drive prices
int b: the budget

Returns
int: the maximum that can be spent, or -1 if it is not possible to buy both items

Input Format
The first line contains three space-separated integers b, n, and m, the budget,
the number of keyboard models and the number of USB drive models.
The second line contains n space-separated integers keyboard[i], the prices of each keyboard model.
The third line contains m space-separated integers drives[i], the prices of the USB drives.

Constraints
1 <= n,m <= 1000
1 <= b <= 10^6
The price of each item is in the inclusive range [1, 10^6].
*/

package main

import "fmt"

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	/*
	 * Write your code here.
	 */
	var greatest int32 = -1

	for _, keyboard := range keyboards {
		for _, drive := range drives {
			var price int32 = keyboard + drive

			if price <= b && price > greatest {
				greatest = price
			} else {
				continue
			}
		}
	}

	return greatest
}

func main() {
	var b, element int32
	var n, m int
	var keyboards, drives []int32 = []int32{}, []int32{}

	fmt.Scanf("%v\n", &b)
	fmt.Scanf("%v\n", &n)
	fmt.Scanf("%v\n", &m)

	for i := 0; i < n; i++ {
		fmt.Scanf("%v\n", &element)
		keyboards = append(keyboards, element)
	}

	for i := 0; i < m; i++ {
		fmt.Scanf("%v\n", &element)
		drives = append(drives, element)
	}

	fmt.Println(getMoneySpent(keyboards, drives, b))
}
