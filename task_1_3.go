package main

import "fmt"

func main() {
	fmt.Println("Task 1_3")
	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(Mutation(in, 3))
}

//Mutation function
func Mutation(in []int, column int) [][]int {
	result := make([][]int, 0)
	counter := 1
	tmp := make([]int, 0)
	for index, el := range in {
		if counter <= column {
			tmp = append(tmp, el)
		}
		counter ++
		if counter > column   {
			counter = 1
			result = append(result, tmp)
			tmp = [] int{}
		} else if index == len(in) - 1{
			dif := column - len(tmp)
			for i :=0; i<dif; i++{
				tmp = append(tmp, 0)
			}
			result = append(result, tmp)
		}
	}
	return result
}
