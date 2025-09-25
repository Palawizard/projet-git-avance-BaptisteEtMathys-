package main

import "fmt"

func somme(a int, b int) int {
	return a + b
}

func main() {
	result := somme(3, 5)
	fmt.Println("The result is :", result)
}
