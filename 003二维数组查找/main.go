package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 12, 13}}
	result := Find(matrix, 3, 4, 10)
	fmt.Printf("%v", result)
}
