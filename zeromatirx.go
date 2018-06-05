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
			nullifyRow(matrix,positionArray[i])
		} else {
			nullifyColumn(matrix,positionArray[i])
		}
	}
	fmt.Println(matrix)
}


 func nullifyRow(matrix [][]int, column int) {
	for i:=0 ; i< len(matrix) ; i++ {
		matrix[i][column]=0
	}
}


func nullifyColumn(matrix [][]int, row int) {
        for i:=0 ; i< len(matrix[0]) ; i++ {
                matrix[row][i]=0
        }
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
