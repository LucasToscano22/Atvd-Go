package main

import "fmt"

var numeros []int
var flag int = 0

func main() {

	for {
		var opcao int

		fmt.Println()
		fmt.Println("------------")
		fmt.Println("Menu")
		fmt.Println("1 - Adicionar número")
		fmt.Println("2 - Listar números")
		fmt.Println("3 - Remover número")
		fmt.Println("4 - Ver estatisticas")
		fmt.Println("5 - Divisão de números")
		fmt.Println("6 - Limpar lista de números")
		fmt.Println("0 - Encerrar")
		fmt.Println("------------")
		fmt.Println()

		fmt.Print("Digite sua opção: ")
		fmt.Scanln(&opcao)

		if opcao == 1 {
			adicionarNumero()
		}

		if opcao == 2 {
			listarNumeros()
		}

		if opcao == 3 {
			removerNumero()
		}

		if opcao == 4 {
			if len(numeros) == 0 {
				fmt.Println("Lista vazia")
				continue
			}

			min, max, media := calcularEstatisticas()

			fmt.Println("Mínimo: ", min)
			fmt.Println("Máximo: ", max)
			fmt.Println("Média: ", media)
		}

		if opcao == 5 {
			dividirNumeros()
		}

		if opcao == 6 {
			limparLista()
		}

		if opcao == flag {
			break
		}

	}
}

func adicionarNumero() {
	var numero int
	fmt.Println("Digite um número para adicionar na lista:")
	fmt.Scanln(&numero)
	numeros = append(numeros, numero)
}

func listarNumeros() {

	if len(numeros) == 0 {
		fmt.Println("Lista está vazia")
		return
	}

	for i := 0; i < len(numeros); i++ {
		fmt.Println(numeros[i])
	}
}

func removerNumero() {

}

func calcularEstatisticas() (int, int, float64){

	// if len(numeros) == 0 {
	// 	return 0, 0, 0
	// }

	min := numeros[0]
	max := numeros[0]
	soma := 0


	for i := 0; i< len(numeros);  i++ {
		if numeros[i] < min {
			min = numeros[i]
		}
		if numeros[i] > max {
			max = numeros[i]
		}
		soma += numeros[i]
	}

	media := float64(soma) / float64(len(numeros))
	
	return min, max, media

}

func dividirNumeros() {

}

func limparLista() {

	if len(numeros) == 0 {
		fmt.Println("Lista já está vazia")
		return
	}

	numeros = []int{}
	fmt.Println("Lista limpa")	
}


