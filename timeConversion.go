/*
Given a time in 12-hour AM/PM format, convert it to military (24-hour) time.

Note:
12:00:00AM on a 12-hour clock is 00:00:00 on a 24-hour clock.
12:00:00PM on a 12-hour clock is 12:00:00 on a 24-hour clock.

Example
s = '12:01:00PM', Return '12:01:00'.
s = '12:01:00AM', Return '00:01:00'.

Function Description
Complete the timeConversion function in the editor below.
It should return a new string representing the input time in 24 hour format.
timeConversion has the following parameter(s):
string s: a time in 12 hour format

Returns
string: the time in 24 hour format

Input Format
A single string s that represents a time in 12-hour clock format
(i.e.: hh:mm:ssAM or hh:mm:ssPM).

Constraints
All input times are valid
*/

package main

import (
	"fmt"
	"strconv"
)

func timeConversion(s string) string {
	// Write your code here
	var time string

	if string(s[8]) == "A" {
		if s[:2] == "12" {
			time = "00" + s[2:8]
		} else {
			time = s[:8]
		}
	} else {
		if s[:2] == "12" {
			time = s[:8]
		} else {
			converted_int, _ := strconv.Atoi(s[:2])
			converted_str := strconv.Itoa(converted_int + 12)
			time = converted_str + s[2:8]
		}
	}

	return time
}

func TimeConversion() {
	var s string

	fmt.Scanf("%s\n", &s)
	fmt.Println(timeConversion(s))
}
