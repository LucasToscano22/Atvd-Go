package main

import "fmt"

var numeros []int

func main() {
	var flag int = 0

	for {
		var opcao int

		fmt.Println()
		fmt.Println("------------")
		fmt.Println("Menu")
		fmt.Println("1 - Adicionar número")
		fmt.Println("2 - Listar números")
		fmt.Println("------------")
		fmt.Println()

		fmt.Print("Digite sua opção: ")
		fmt.Scanln(&opcao)

		if opcao == 1 {
			fmt.Println("Teste, digitou o numero 1")
		}

		if opcao == 2 {
			fmt.Println("Teste, digitou o numero 2")
		}

		if opcao == flag {
			break
		}

	}
}

