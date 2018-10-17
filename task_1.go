package main

import "fmt"

func main() {
	fmt.Println("Hello")
	in := []int{9, 1, 9, 5, 4, 4, 2, 1, 5, 4, 8, 8, 4, 3, 6, 9, 5, 7, 5}
	arr := InsertUn(in)
	fmt.Println(arr)
}

// InsertUn function to remove duplicates
func InsertUn(in []int) []int {
	m := make(map[int]int)
	result := []int{}
	for _, item := range in {
		if _, found := m[item]; !found {
			result = append(result, item)
			m[item] = item
		}
	}

	return result
}
