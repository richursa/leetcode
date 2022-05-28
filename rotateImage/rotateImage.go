package main

func rotate(matrix [][]int) {
	recursiveRotate(matrix, len(matrix), 0)
}

func recursiveRotate(matrix [][]int, length int, offset int) {
	if length <= 1 {
		return
	}
	for i := 0; i < length-1; i++ {
		matrix[offset][i+offset], matrix[i+offset][length-1+offset], matrix[length-1+offset][length-1-i+offset], matrix[length-1-i+offset][0+offset] = matrix[length-1-i+offset][0+offset], matrix[0+offset][i+offset], matrix[i+offset][length-1+offset], matrix[length-1+offset][length-1-i+offset]
	}
	recursiveRotate(matrix, length-2, offset+1)
}
