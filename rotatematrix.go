package main

import(
"fmt"
)

func MatrixRotate(matrix [][]int) {
	fmt.Println(matrix)
	n:= len(matrix)
	for i:=0; i<n/2 ;i++ {
		m := len(matrix[i])
		for j:=i ; j<n-i-1; j++ {
			temp := matrix[i][j]
			
			matrix[i][j] = matrix[n-i-1][j]
			matrix[n-i-1][j]=matrix[n-i-1][m-j-1]
			matrix[n-i-1][m-j-1] = matrix[i][m-j-1];
			matrix[i][m-j-1] = temp
			
		}
	}
	fmt.Println(matrix)
}

func main() {
	matrix := [][]int{}
	matrix1 := []int{1,2,3}
	matrix2 := []int{33,66,99}
	matrix3 := []int{16,17,18}
	matrix = append(matrix,matrix1)
	matrix = append(matrix,matrix2)
	matrix = append(matrix,matrix3)
	MatrixRotate(matrix)
}
				
