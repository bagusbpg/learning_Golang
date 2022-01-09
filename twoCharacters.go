/*
Given a string, remove characters until the string is made up of
any two alternating characters. When you choose a character to remove,
all instances of that character must be removed. Determine the longest
string possible that contains just two alternating letters.

Example
s = 'abaacdabd'
Delete a, to leave bcdbd. Now, remove the character c to leave the valid
string bdbd with a length of 4. Removing either b or d at any point
would not result in a valid string. Return 4.
Given a string s, convert it to the longest possible string t made up only
of alternating characters. Return the length of string t. If no string t
can be formed, return 0.

Function Description
Complete the alternate function in the editor below.
alternate has the following parameter(s):
string s: a string

Returns.
int: the length of the longest valid string, or 0 if there are none,

Input Format
The first line contains a single integer that denotes the length of s.
The second line contains string s.

Constraints
1 <= length of s <= 1000
s[i] E ascii[a-z]
*/

package main

import (
	"fmt"
	"regexp"
)

func alternate(s string) int32 {
	// Write your code here
	var mapChar map[rune]int = map[rune]int{}
	var longest int32 = 0

	for _, character := range s {
		if _, exist := mapChar[character]; !exist {
			mapChar[character] = 1
		}
	}

	var newMapChar map[int]rune = map[int]rune{}
	var index int = 0

	for character := range mapChar {
		newMapChar[index] = character
		index++
	}

	var regexes []string = []string{}

	for i := 0; i < len(newMapChar)-1; i++ {
		for j := i + 1; j < len(newMapChar); j++ {
			var regex string = "[^" + string(newMapChar[i]) + string(newMapChar[j]) + "]"
			regexes = append(regexes, regex)
		}
	}

	for _, regex := range regexes {
		re := regexp.MustCompile(regex)
		newstr := re.ReplaceAll([]byte(s), []byte(""))

		if len(newstr)%2 != 0 && newstr[0] != newstr[len(newstr)-1] {
			continue
		}

		var flag bool = true

		for i := 2; i <= len(newstr)-2 && flag; i += 2 {
			if newstr[i] != newstr[0] || newstr[i+1] != newstr[1] {
				flag = false
			}
		}

		if flag {
			if int32(len(newstr)) > longest {
				longest = int32(len(newstr))
			}
		}
	}

	return longest
}

func main() {
	var s string

	fmt.Scanf("%v\n", &s)

	fmt.Println(alternate(s))
}
