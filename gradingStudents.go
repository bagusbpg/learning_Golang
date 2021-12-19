/*
HackerLand University has the following grading policy:
1. Every student receives a grade in the inclusive range from 0 to 100.
2. Any grade less than 40 is a failing grade.

Sam is a professor at the university and likes to round each student's grade
according to these rules:
1. If the difference between the grade and the next multiple of 5 is less than 3,
round up to the next multiple of 5.
2. If the value of grade is less than 38, no rounding occurs as the result will
still be a failing grade.

Examples
grade = 84 round to (85 - 84 is less than 3)
grade = 29 do not round (result is less than 40)
grade = 57 do not round (60 - 57 is 3 or higher)

Given the initial value of grade for each of Sam's n students,
write code to automate the rounding process.

Function Description
Complete the function gradingStudents in the editor below.
gradingStudents has the following parameter(s):
int grades[n]: the grades before rounding

Returns
int[n]: the grades after rounding as appropriate

Input Format
The first line contains a single integer, n, the number of students.
Each line i of the n subsequent lines contains a single integer, grades[i].

Constraints
1 <= n <= 60
0 <= grades[i] <= 100
*/

package main

import "fmt"

func gradingStudents(grades []int32) []int32 {
	// Write your code here
	var result []int32 = []int32{}

	for _, grade := range grades {
		if grade < 38 {
			result = append(result, grade)
		} else {
			if ((grade + 2) / 5) > (grade / 5) {
				var new_score int32 = int32(((grade + 2) / 5)) * 5
				result = append(result, new_score)
			} else {
				result = append(result, grade)
			}
		}
	}

	return result
}

func main() {
	var n int
	var element int32
	var ar []int32 = []int32{}

	fmt.Scanf("%v\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v\n", &element)
		ar = append(ar, element)
	}

	for _, value := range gradingStudents(ar) {
		fmt.Println(value)
	}
}
