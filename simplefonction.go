package main

import "fmt"

func addition(a int, b int) int {
	return a + b
}

func main() {
	resultat := addition(3, 5)
	fmt.Println("Le résultat est :", resultat)
}
