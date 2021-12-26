/*
When a contiguous block of text is selected in a PDF viewer,
the selection is highlighted with a blue rectangle.
In this PDF viewer, each word is highlighted independently.

There is a list of 26 character heights aligned by index to their letters.
For example, 'a' is at index 0 and 'z' is at index 25.
There will also be a string. Using the letter heights given,
determine the area of the rectangle highlight in mmsq assuming
all letters are 1mm wide.

Example
h = [1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 1, 1, 5, 5, 1, 5, 2, 5, 5, 5, 5, 5, 5]
    a   b  c  d  e  f  g  h  i  j  k  l  m  n  o  p  q  r  s  t  u  v  w  x  y  z
word = 'torn'
The heights are t = 2, o = 1, r = 1 and n = 1.
The tallest letter is 2 high and there are 4 letters.
The hightlighted area will 2 * 4 = 8mmsq be so the answer is 8.

Function Description
Complete the designerPdfViewer function in the editor below.
designerPdfViewer has the following parameter(s):
int h[26]: the heights of each letter
string word: a string

Returns
int: the size of the highlighted area

Input Format
The first line contains 26 space-separated integers describing
the respective heights of each consecutive lowercase English letter, ascii[a-z].
The second line contains a single word consisting of lowercase English alphabetic letters.

Constraints
1 <= h[?] <= 7, where ? is an English lowercase letter.
word contains no more than 10 letters.
*/

package main

import "fmt"

func designerPdfViewer(h []int32, word string) int32 {
	// Write your code here
	var highest int32 = 0

	for _, character := range word {
		if h[character-97] > highest {
			highest = h[character-97]
		}
	}

	return highest * int32(len(word))
}

func main() {
	var ar []int32 = []int32{}
	var element int32
	var word string

	for i := 0; i < 26; i++ {
		fmt.Scanf("%v\n", &element)
		ar = append(ar, element)
	}

	fmt.Scanf("%v\n", &word)

	fmt.Println(designerPdfViewer(ar, word))
}
