package main

import(
"fmt"
)

func MatrixRotate(matrix [][]int) {
	fmt.Println(matrix)
	n:= len(matrix)
	fmt.Println(n/2)
	m:=len(matrix[0])
	if m!=n {
		fmt.Println("Not a square matrix")
		return
	}
	for layer:=0; layer<n/2 ;layer++ {
		last:= n-1-layer
		for j:=layer ; j<n-layer-1; j++ {
			offset := j-layer
			temp := matrix[layer][j]	
			matrix[layer][j] = matrix[last - offset][layer]
			matrix[last-offset][layer]=matrix[n-layer-1][last-offset]
			matrix[n-layer-1][last-offset] = matrix[j][n-layer-1];
			matrix[j][n-layer-1] = temp
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
				
