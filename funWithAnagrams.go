/*
Two strings ara anagrams if they are permutations of each other. In other words,
both strings have the same size and the same characters. For example, "aaagmnrs"
is an anagram of "anagrams". Given an array of strings, remove each string that is
an anagram of an earlier string, then return the remaining array in sorted order.

Example
str = ['code', 'doce', 'ecod', 'framer', 'frame']

'code' and 'doce' ara anagrams. Remove 'doce' from the array and keep the first
occurrence 'code' in the array.
'code' and 'ecod' ara anagrams. Remoce 'ecod' from the array and keep the first
occurrence 'code' in the array.
'code' and 'framer' are not anagrams. Keep both strings in the array.
'framer' and 'frame' are not anagrams due to the extra 'r' in 'framer'. Keep both
strings in the array.
Order the remaining strings in ascending order : ['code', 'frame', 'framer'].

Function Description
Complete the function funWithAnagrams in the editor below.
funWithAnagrams has the following parameters:
string text[n]: an array of strings

Returns
string[m]: an array of the remaining strings in ascending alphabetical order.

Constraints
0 <= n <= 1000
0 <= m <= n
1 <= length of text[i] <= 1000
Each string text[i] is made up of characters in the range ascii[a-z]
*/

package main

import (
	"fmt"
	"sort"
)

func funWithAnagrams(texts []string) []string {
	// Write your code here
	var newtexts []string = []string{}

	for _, text := range texts {
		var intsOfRune []int = []int{}

		for _, value := range text {
			intsOfRune = append(intsOfRune, int(value))
		}

		sort.Ints(intsOfRune)
		var key string = ""

		for _, value := range intsOfRune {
			key += fmt.Sprint(value)
		}

		newtexts = append(newtexts, key)
	}

	var maptexts map[string]int = map[string]int{}
	var result []string = []string{}

	for i := 0; i < len(newtexts); i++ {
		if _, exist := maptexts[newtexts[i]]; !exist {
			maptexts[newtexts[i]] = 1
			result = append(result, texts[i])
		}
	}

	sort.Strings(result)

	return result
}

func main() {
	var n int
	var element string
	var ar []string = []string{}

	fmt.Scanf("%v\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v\n", &element)
		ar = append(ar, element)
	}

	fmt.Println(funWithAnagrams(ar))
}
