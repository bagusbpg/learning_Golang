/*
You have two strings of lowercase English letters. You can perform
two types of operations on the first string:
1. Append a lowercase English letter to the end of the string.
2. Delete the last character of the string. Performing this operation
on an empty string results in an empty string.

Given an integer, k, and two strings, s and t, determine whether or not
you can convert s to t by performing exactly k of the above operations on s.
If it's possible, print Yes. Otherwise, print No.

Example
s = [a, b, c]
t = [d, e, f]
k = 6
To convert s to t, we first delete all of the characters in 3 moves.
Next we add each of the characters of t in order. On the 6th move, you will
have the matching string. Return Yes.
If there were more moves available, they could have been eliminated by performing
multiple deletions on an empty string. If there were fewer than 6 moves,
we would not have succeeded in creating the new string.

Function Description
Complete the appendAndDelete function in the editor below.
It should return a string, either Yes or No.
appendAndDelete has the following parameter(s):
string s: the initial string
string t: the desired string
int k: the exact number of operations that must be performed

Returns
string: either Yes or No

Input Format
The first line contains a string s, the initial string.
The second line contains a string t, the desired final string.
The third line contains an integer k, the number of operations.

Constraints
1 <= |s|, |t|, k <= 100
s and t consist of lowercase English letters, ascii[a-z].

Note, I did not understand the test case 2:
s = "aaaaaaaaaa"
t = "aaaaa"
k = 7
It returns Yes. But if we alter s to
s = "aaaaaaaaa" (one letter less)
It returns No! It does not make sense.

I also find that test case 5 is hard to understand.
s = "y"
t = "yu"
k = 2
It returns No!

Test case 10 is also somewhat strange
s = "abcd"
t = "abcdert"
k = 10
It returns No! But if we alter s to
s = "abcde" (one letter more)
It returns Yes!
*/

package main

import "fmt"

func appendAndDelete(s string, t string, k int32) string {
	// Write your code here
	var index = 0
	var flag bool = true

	for index < len(s) && index < len(t) && flag {
		if s[index] == t[index] {
			index++
		} else {
			flag = false
		}
	}

	if len(s)+len(t)-2*index > int(k) {
		return "No"
	} else if len(s)+len(t)-2*index == int(k) {
		return "Yes"
	} else if len(s)+len(t) <= int(k) {
		return "Yes"
	} else if (len(s)+len(t)-int(k))%2 == 0 {
		return "Yes"
	} else {
		return "No"
	}
}

func main() {
	var s, t string
	var k int32

	fmt.Scanf("%v\n", &s)
	fmt.Scanf("%v\n", &t)
	fmt.Scanf("%v\n", &k)

	fmt.Println(appendAndDelete(s, t, k))
}
