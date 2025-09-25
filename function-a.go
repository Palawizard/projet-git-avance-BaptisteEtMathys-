package main

import "fmt"

// functionA prints "Function A" to the console
func functionA(A string) {
	fmt.Println("Function A")
	fmt.Println("You have entered", A, "as an argument")
}

func main() {
	functionA("a")
}
