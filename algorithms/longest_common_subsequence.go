// Implementation of the Longest Common Subsequence DP algorithm in Go

package main

import "fmt"

func createByteMatrix(height, width int) [][][]byte {
	matrix := make([][][]byte, height)
	for index := range matrix {
		row := make([][]byte, width)
		matrix[index] = row
	}

	return matrix
}

func getLongestAdjacentSequence(matrix [][][]byte, i, j int) []byte {
	l := matrix[i][j-1]
	t := matrix[i-1][j]

	if len(l) > len(t) {
		return l
	}
	return t
}

func LongestCommonSubsequence(s1 string, s2 string) []byte {
	if s1 == "" || s2 == "" {
		return []byte{}
	}

	matrix := createByteMatrix(len(s1)+1, len(s2)+1)

	for i := 1; i < len(s1)+1; i++ {
		for j := 1; j < len(s2)+1; j++ {

			if s1[i-1] == s2[j-1] {
				newRow := make([]byte, len(matrix[i-1][j-1]))
				copy(newRow, matrix[i-1][j-1])
				newRow = append(newRow, s1[i-1])
				matrix[i][j] = newRow
			} else {
				row := getLongestAdjacentSequence(matrix, i, j)
				matrix[i][j] = row
			}
		}
	}

	return matrix[len(s1)][len(s2)]
}

func main() {
	s1 := "Hello World"
	s2 := "Work"

	fmt.Printf("%s\n", LongestCommonSubsequence(s1, s2))
}
