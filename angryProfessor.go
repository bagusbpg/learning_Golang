/*
A Discrete Mathematics professor has a class of students.
Frustrated with their lack of discipline, the professor decides to cancel class
if fewer than some number of students are present when class starts.
Arrival times go from on time (arrivalTime <= 0) to arrived late (arrivalTime > 0).

Given the arrival time of each student and a threshhold number of attendees,
determine if the class is cancelled.

Example
n = 5
k = 3
a = [-2, -1, 0, 1, 2]
The first 3 students arrived on. The last 2 were late. The threshold is 3 students,
so class will go on. Return NO.
Note: Non-positive arrival times (a[i] <= 0) indicate the student arrived early or on time;
positive arrival times (a[i] > 0) indicate the student arrived a[i] minutes late.

Function Description
Complete the angryProfessor function in the editor below.
It must return YES if class is cancelled, or NO otherwise.
angryProfessor has the following parameter(s):
int k: the threshold number of students
int a[n]: the arrival times of the n students

Returns
string: either YES or NO

Input Format
The first line has two space-separated integers, n and k,
the number of students (size of ar) and the cancellation threshold.
The second line contains n space-separated integers (a[1], a[2], ..., a[n])
that describe the arrival times for each student.

Constraints
1 <= n <= 1000
1 <= k <= n
-100 <= a[i] <= 100, where i E [0, ..., n-1]
*/

package main

import "fmt"

func angryProfessor(k int32, ar []int32) string {
	// Write your code here
	var ontime int32 = 0

	for _, arrival := range ar {
		if arrival <= 0 {
			ontime++
		}
	}

	if ontime >= k {
		return "NO"
	} else {
		return "YES"
	}
}

func AngryProfessor() {
	var n int
	var element, k int32
	var ar []int32 = []int32{}

	fmt.Scanf("%v\n", &n)
	fmt.Scanf("%v\n", &k)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v\n", &element)
		ar = append(ar, element)
	}

	fmt.Println(angryProfessor(k, ar))
}
