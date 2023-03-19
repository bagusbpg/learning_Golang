/*
Given s and x, determine the zero-based index of the first occurence of x in s.
String s consists of lowercase letters in the range ascii[a-z].
String x consists of lowercase letters and may also contain a single wildcard
character, *, that represents any one character.

Example
s = "xabcdey"
x = "ab*de"
The first match is at index 1.
s	  x a b c d e y
x       a b * d e
index 0 1 2 3 4 5 6

Function Description
Complete the function firstOccurene in the editor below. The function must return
an integer denoting the zero-based index of the first occurence of string x in s.
If x is not in s, return -1 instead.
firstOccurrence has the following parameter(s):
string s: a string of lowercase letters
string x: a string of lowercase letters which may contain 1 instance of the wildcard
character *

Returns
int: the index of the first occurrence, or -1 if x does not occur in s.

Constraints
1 <= length of s <= 5 * 10^5
1 <= length of x <= 1000
*/

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func firstOccurrence(s string, x string) int32 {
	// Write your code here
	var new_x string = strings.ReplaceAll(x, "*", ".")

	pattern := regexp.MustCompile(new_x)
	s_byte := []byte(s)
	loc := pattern.FindIndex(s_byte)

	if len(loc) > 0 {
		return int32(loc[0])
	} else {
		return -1
	}
}

func FirstOccurrence() {
	var s, x string

	fmt.Scanf("%v\n", &s)
	fmt.Scanf("%v\n", &x)

	fmt.Println(firstOccurrence(s, x))
}
