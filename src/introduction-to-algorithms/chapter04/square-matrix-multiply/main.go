package main

import (
	"fmt"
)

const (
	m = 3
	n = 4
)

func squareMatrixMultiply(matA [m][n]int, matB [n][m]int) (matC [m][m]int) {
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < n; k++ {
				matC[i][j] += matA[i][k] * matB[k][j]
			}
		}
	}
	return
}

func main() {
	matA := [m][n]int{{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}}
	matB := [n][m]int{{5, 6, 7}, {5, 6, 7}, {5, 6, 7}, {5, 6, 7}}
	matC := squareMatrixMultiply(matA, matB)
	fmt.Println(matA)
	fmt.Println(matB)
	fmt.Println(matC)
}
