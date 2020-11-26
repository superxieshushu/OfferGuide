package main

//二维数组查找
//二维数组matrix中，每一行都从左到右递增排序，
//每一列都从上到下递增排序
func Find(matrix [][]int, rows int, columns int, number int) (isFind bool) {

	//边界情况
	if matrix == nil || rows <= 0 || columns <= 0 {
		return
	}

	//主逻辑
	row := 0
	column := columns - 1
	//循环的条件是啥？变化的条件是啥？
	for row < rows && column >= 0 {
		if matrix[row][column] == number {
			return true
		}
		if matrix[row][column] < number {
			row++
		} else {
			column--
		}
	}
	return
}
