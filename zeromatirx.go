package main

import (
	"fmt"
)

func zeromatrix(matrix [][]int) {
	var positionArray []int
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				positionArray = append(positionArray, i)
				positionArray = append(positionArray, j)
			}
		}
	}

	fmt.Println(positionArray)
	for i := 0; i < len(positionArray); i++ {
		if i%2 != 0 {
		} else {
			i1 := i + 1
			for a := 0; a < m; a++ {
				k := positionArray[i1]
				matrix[a][k] = 0
			}
			for a := 0; a < n; a++ {
				k := positionArray[i]
				matrix[k][a] = 0
			}
		}
	}
	fmt.Println(matrix)
}

func main() {
	var matrix [][]int
	matrix1 := []int{1, 0, 3}
	matrix2 := []int{0, 1, 1}
	matrix3 := []int{234, 1432, 2354}
	matrix4 := []int{3, 2, 1}
	matrix = append(matrix, matrix1)
	matrix = append(matrix, matrix2)
	matrix = append(matrix, matrix3)
	matrix = append(matrix, matrix4)
	zeromatrix(matrix)
}
