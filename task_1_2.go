package main

import "fmt"

func main() {
	fmt.Println("Task 1_2")
	irregularMatrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11},
		{12, 13, 14, 15},
		{16, 17, 18, 19, 20}}
	slice := Flatten(irregularMatrix)
	fmt.Printf("1x%d: %v\n", len(slice), slice)
}

// Flatten function
func Flatten(in [][]int) []int {
	result := []int{}
	for _, t := range in {
		for _, it := range t {
			result = append(result, it)
		}
	}
	return result
}
