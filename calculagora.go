package main

import (
	"fmt"
)

func main() {
	var n1, n2 float64
	var op string

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&n1)
	fmt.Print("Digite o operador (+, -, *, /): ")
	fmt.Scanln(&op)
	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&n2)

	switch op {
	case "+":
		fmt.Printf("Resultado: %.2f\n", n1+n2)
	case "-":
		fmt.Printf("Resultado: %.2f\n", n1-n2)
	case "*":
		fmt.Printf("Resultado: %.2f\n", n1*n2)
	case "/":
		if n2 != 0 {
			fmt.Printf("Resultado: %.2f\n", n1/n2)
		} else {
			fmt.Println("Erro: Divisão por zero!")
		}
	default:
		fmt.Println("Operador inválido!")
	}
}
