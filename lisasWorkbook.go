/*
Lisa just got a new math workbook. A workbook contains exercise problems,
grouped into chapters. Lisa believes a problem to be special if its index
(within a chapter) is the same as the page number where it's located.
The format of Lisa's book is as follows:
1. There are n chapters in Lisa's workbook, numbered from 1 to n.
2. The ith chapter has arr[i] problems, numbered from 1 to arr[i].
3. Each page can hold up to k problems. Only a chapter's last page
of exercises may contain fewer than k problems.
4. Each new chapter starts on a new page, so a page will never contain
problems from more than one chapter.
5. The page number indexing starts at 1.
Given the details for Lisa's workbook, can you count its number of
special problems?

Example
arr = [4, 2]
k = 3
Lisa's workbook contains arr[1] = 4 problems for chapter 1, and
arr[2] problems for chapter 2. Each page can hold k = 3 problems.
The first page will hold 3 problems for chapter 1. Problem 1 is on page 1,
so it is special. Page 2 contains only Chapter 1, Problem 4, so no special
problem is on page 2. Chapter 2 problems start on page 3 and there are
2 problems. Since there is no problem 3 on page 3, there is no special
problem on that page either. There is 1 special problem in her workbook.
Note: See the diagram in the Explanation section for more details.

Function Description
Complete the workbook function in the editor below.
workbook has the following parameter(s):
int n: the number of chapters
int k: the maximum number of problems per page
int arr[n]: the number of problems in each chapter

Returns
int: the number of special problems in the workbook

Input Format
The first line contains two integers n and k, the number of chapters
and the maximum number of problems per page.
The second line contains n space-separated integers arr[i] where arr[i]
denotes the number of problems in the ith chapter.

Constraints
1 <= n, k, arr[i] <= 100
*/

package main

import "fmt"

func workbook(n int32, k int32, arr []int32) int32 {
	// Write your code here
	var problems []int32 = []int32{}
	var count int32 = 0

	for i := 0; i < int(n); i++ {
		for j := int32(1); j <= arr[i]; j++ {
			problems = append(problems, j)
		}
		if arr[i]%k != 0 {
			for l := 0; l < int(k-arr[i]%k); l++ {
				problems = append(problems, 0)
			}
		}
	}

	for i := 0; i < len(problems); i++ {
		if problems[i] == int32(i)/k+1 {
			count++
		}
	}

	return count
}

func Workbook() {
	var n, k, element int32
	var ar []int32 = []int32{}

	fmt.Scanf("%v\n", &n)
	fmt.Scanf("%v\n", &k)
	for i := int32(0); i < n; i++ {
		fmt.Scanf("%v\n", &element)
		ar = append(ar, element)
	}

	fmt.Println(workbook(n, k, ar))
}
