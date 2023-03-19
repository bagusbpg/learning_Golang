/*
There are a number of people who will be attending ACM-ICPC World Finals.
Each of them may be well versed in a number of topics. Given a list of topics
known by each attendee, presented as binary strings, determine the maximum number
of topics a 2-person team can know. Each subject has a column in the binary string,
and a '1' means the subject is known while '0' means it is not. Also determine
the number of teams that know the maximum number of topics. Return an integer array
with two elements. The first is the maximum number of topics known, and the second
is the number of teams that know that number of topics.

Example
n = 3
topics = ['10101', '1110', '00010']
The attendee data is aligned for clarity below:

10101
11110
00010

These are all possible teams that can be formed:

Members Subjects
(1,2)   [1,2,3,4,5]
(1,3)   [1,3,4,5]
(2,3)   [1,2,3,4]

In this case, the first team will know all 5 subjects.
They are the only team that can be created that knows that many subjects,
so [5, 1] is returned.

Function Description
Complete the acmTeam function in the editor below.
acmTeam has the following parameter(s):
string topic: a string of binary digits

Returns
int[2]: the maximum topics and the number of teams that know that many topics

Input Format
The first line contains an integer n, the number of attendees.
Each of the next n lines contains a binary string of length m.

Constraints
2 <= n <= 500
1 <= m <= 500
*/

package main

import "fmt"

func acmTeam(topic []string) []int32 {
	// Write your code here
	var length int = len(topic[0])
	var mapTopic map[int32]int32 = map[int32]int32{}
	var greatest int32 = 0

	for i := 0; i < len(topic)-1; i++ {
		for j := i + 1; j < len(topic); j++ {
			var count int32 = 0
			for k := 0; k < length; k++ {
				if topic[i][k] == 49 || topic[j][k] == 49 {
					count++
				}
			}

			if _, exist := mapTopic[count]; exist {
				mapTopic[count]++
			} else {
				mapTopic[count] = 1
			}
		}
	}

	for key := range mapTopic {
		if key > greatest {
			greatest = key
		}
	}

	return []int32{greatest, mapTopic[greatest]}
}

func ACMTeam() {
	var n int
	var element string
	var ar []string = []string{}

	fmt.Scanf("%v\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v\n", &element)
		ar = append(ar, element)
	}

	fmt.Println(acmTeam(ar))
}
