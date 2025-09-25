package main

import "fmt"

func somme(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	result := somme(3, 5, 7, 9, 11)
	fmt.Println("The result is :", result)
}
