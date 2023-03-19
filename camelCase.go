/*
There is a sequence of words in CamelCase as a string of letters, s,
having the following properties:
1. It is a concatenation of one or more words consisting of English letters.
2. All letters in the first word are lowercase.
3. For each of the subsequent words, the first letter is uppercase and rest of the letters are lowercase.

Given s, determine the number of words in s.

Example
s = oneTwoThree
There are 3 words in the string: 'one', 'Two', 'Three'.

Function Description
Complete the camelcase function in the editor below.
camelcase has the following parameter(s):
string s: the string to analyze

Returns
int: the number of words in s.

Input Format
A single line containing string s.

Constraints
1 <= length of s <= 10^5
*/

package main

import (
	"fmt"
	"unicode"
)

func camelcase(s string) int32 {
	// Write your code here
	var count int32 = 1

	for _, character := range s {
		if unicode.IsUpper(character) {
			count++
		}
	}

	return count
}

func Camelcase() {
	var s string

	fmt.Scanf("%v\n", &s)

	fmt.Println(camelcase(s))
}
