/*
There is a collection of rocks where each rock has various minerals
embeded in it. Each type of mineral is designated by a lowercase letter
in the range ascii[a-z]. There may be multiple occurrences of a mineral
in a rock. A mineral is called a gemstone if it occurs at least once in
each of the rocks in the collection.
Given a list of minerals embedded in each of the rocks, display the number
of types of gemstones in the collection.

Example
arr = ['abc', 'abc', 'bc']
The minerals b and c appear in each rock, so there are 2 gemstones.

Function Description
Complete the gemstones function in the editor below.
gemstones has the following parameter(s):
string arr[n]: an array of strings

Returns
int: the number of gemstones found

Input Format
The first line consists of an integer n, the size of arr.
Each of the next n lines contains a string arr[i] where each letter
represents an occurence of a mineral in the current rock.

Constraints
1 <= n <= 100
1 <= |arr[i]| <= 100
Each composition arr[i] consists of only lower-case Latin letters
('a'-'z').

Sample Input
STDIN       Function
-----       --------
3           arr[] size n = 3
abcdde      arr = ['abcdde', 'baccd', 'eeabg']
baccd
eeabg

Sample Output
2

Explanation
Only a and b occur in every rock.
*/

package main

import "fmt"

func gemstones(arr []string) int32 {
	// Write your code here
	if len(arr) == 1 {
		return 1
	}

	mapOfChar := make(map[rune]int)
	for _, char := range arr[0] {
		if _, exist := mapOfChar[char]; !exist {
			mapOfChar[char] = 1
		}
	}

	for _, s := range arr[1:] {
		uniqueChar := make(map[rune]struct{})
		for _, char := range s {
			if _, exist := uniqueChar[char]; exist {
				continue
			}

			uniqueChar[char] = struct{}{}

			if _, exist := mapOfChar[char]; exist {
				mapOfChar[char] += 1
			}
		}
	}

	var numberOfGemstone int32 = 0
	for _, count := range mapOfChar {
		if count == len(arr) {
			numberOfGemstone += 1
		}
	}

	return numberOfGemstone
}

func main() {
	var n int
	var s string
	var ar []string

	fmt.Scanf("%v\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v\n", &s)
		ar = append(ar, s)
	}

	fmt.Println(gemstones(ar))
}
