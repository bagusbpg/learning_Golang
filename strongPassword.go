/*
Louise joined a social networking site to stay in touch with her friends.
The signup page required her to input a name and a password. However,
the password must be strong. The website considers a password to be strong
if it satisfies the following criteria:
1. Its length is at least 6.
2. It contains at least one digit.
3. It contains at least one lowercase English character.
4. It contains at least one uppercase English character.
5. It contains at least one special character. The special characters are:
!@#$%^&*()-+
She typed a random string of length n in the password field but wasn't sure
if it was strong. Given the string she typed, can you find the minimum number
of characters she must add to make her password strong?
Note: Here's the set of types of characters in a form you can paste in your
solution:

numbers = "0123456789"
lower_case = "abcdefghijklmnopqrstuvwxyz"
upper_case = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
special_characters = "!@#$%^&*()-+"

Example
password = '2bbbb'
This password is 5 characters long and is missing an uppercase and a special
character. The minimum number of characters to add is 2.

password = '2bb#A'
This password is 5 characters long and has at least one of each character type.
The minimum number of characters to add is 1.

Function Description
Complete the minimumNumber function in the editor below.
minimumNumber has the following parameters:
int n: the length of the password
string password: the password to test

Returns
int: the minimum number of characters to add

Input Format
The first line contains an integer n, the length of the password.
The second line contains the password string. Each character is
either a lowercase/uppercase English alphabet, a digit, or a special
character.

Constraints
1 <= n <= 100
All characters in password are in [a-z], [A-Z], [0-9], or [!@#$%^&*()-+ ].
*/

package main

import (
	"fmt"
	"regexp"
)

func minimumNumber(n int32, password string) int32 {
	// Return the minimum number of characters to make the password strong
	var count int32 = 0

	if matched, _ := regexp.Match(`[0-9]`, []byte(password)); !matched {
		count++
	}

	if matched, _ := regexp.Match(`[a-z]`, []byte(password)); !matched {
		count++
	}

	if matched, _ := regexp.Match(`[A-Z]`, []byte(password)); !matched {
		count++
	}

	if matched, _ := regexp.Match(`[!@#$%^&*()\-+]`, []byte(password)); !matched {
		count++
	}

	if n+count > 6 {
		return count
	} else {
		return 6 - n
	}
}

func main() {
	var s string

	fmt.Scanf("%v\n", &s)

	fmt.Println(minimumNumber(int32(len(s)), s))
}
