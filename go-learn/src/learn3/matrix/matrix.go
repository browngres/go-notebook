package main

import (
	"fmt"
	"math/rand"
)

type Matrix struct {
	row   int
	col   int
	array *([][]int)
}

func (m *Matrix) printMatrix() {
	// 打印该矩阵
	for _, v := range *(m.array) {
		fmt.Println(v)
	}
}

func newMatrix(row int, col int) *[][]int {
	// 初始化矩阵
	var array [][]int = make([][]int, row)
	for i := 0; i < row; i++ {
		a := make([]int, col)
		for j := 0; j < col; j++ {
			a[j] = rand.Intn(100)
		}
		array[i] = a
	}
	return &array
}

func (m *Matrix) transpose() {
	// 矩阵转置。两两交换，不用新内存
	// 要求方阵
	for i := 0; i < m.row; i++ {
		for j := i + 1; j < m.col; j++ {
			// 交换
			(*(m.array))[i][j], (*(m.array))[j][i] = (*(m.array))[j][i], (*(m.array))[i][j]
		}
	}
}

func main() {
	row := 4
	col := 4
	array := newMatrix(row, col)

	var m1 Matrix
	m1.row = row
	m1.col = col
	m1.array = array
	m1.printMatrix()
	fmt.Println("============转置后==============")
	m1.transpose()
	m1.printMatrix()

}
