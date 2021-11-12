package main

import "fmt"

func main() {
	var (
		numero int
	)

	fmt.Println("É Par ou Impar ?")
	fmt.Print("Digite um número: ")
	fmt.Scan(&numero)

	if numero%2 == 0 {
		fmt.Println("Par")
	} else {
		fmt.Println("Impar")
	}
}
