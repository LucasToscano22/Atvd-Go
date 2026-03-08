package main

import (
	"fmt"
	"errors"
)

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
		fmt.Println("7 - Listar números pares")
		fmt.Println("0 - Encerrar")
		fmt.Println("------------")
		fmt.Println()

		fmt.Print("Digite sua opção: ")
		fmt.Scanln(&opcao)

		if opcao == 1 {
			err := adicionarNumero()
			if err != nil {
				fmt.Println("Erro:", err)
			}
		}

		if opcao == 2 {
			err := listarNumeros()
			if err != nil {
				fmt.Println("Erro:", err)
			}
		}

		if opcao == 3 {
			err := removerNumero()
			if err != nil {
				fmt.Println("Erro:", err)
			}
		}

		if opcao == 4 {
			min, max, media, err := calcularEstatisticas()

			if err != nil {
				fmt.Println("Erro:", err)
				continue
			}

			fmt.Println("Mínimo: ", min)
			fmt.Println("Máximo: ", max)
			fmt.Println("Média: ", media)
		}

		if opcao == 5 {
			err := dividirNumeros()
			if err != nil {
				fmt.Println("Erro:", err)
			}
		}
		
		if opcao == 6 {
			err := limparLista()
			if err != nil {
				fmt.Println("Erro:", err)
			}
		}

		if opcao == 7 {
			err := listarNumerosPares()
			if err != nil {
				fmt.Println("Erro:", err)
			}
		}

		if opcao == flag {
			break
		}

	}
}

func adicionarNumero() error {
	var numero int
	fmt.Println("Digite um número para adicionar na lista:")
	fmt.Scanln(&numero)

	if numero < 0 {
		return errors.New("Apenas números positivos são aceitos")
	}
	numeros = append(numeros, numero)
	return nil
}

func listarNumeros() error {

	if len(numeros) == 0 {
		return errors.New("Lista está vazia")
	}

	for i := 0; i < len(numeros); i++ {
		fmt.Println("Indice: ", i, "Valor: ", numeros[i])
	}

	return nil
}

func removerNumero() error {
	
	if len(numeros) == 0 {
		return errors.New("Lista está vazia")
	}

	fmt.Println("Digite um indice para remover o elemento da lista:")
	var i int
	fmt.Scanln(&i)

	if i < 0  || i >= len(numeros) {
		return errors.New("Indice inválido")
	}

	numeros = append(numeros[:i], numeros[i+1:]...)

	fmt.Println("Número removido da lista")
	return nil
}

func calcularEstatisticas() (int, int, float64, error){

	if len(numeros) == 0 {
		return 0, 0, 0, errors.New("Lista está vazia")
	}

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
	
	return min, max, media, nil

}

func dividir(a int, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("Não é possivel dividir por zero")
	}

	resultadoDivisao := float64(a)/float64(b)
	return resultadoDivisao, nil
}

func dividirNumeros() error {

	var a, b int

	fmt.Println("Digite o primeiro número:")
	fmt.Scanln(&a)
	fmt.Println("Digite o segundo número:")
	fmt.Scanln(&b)
	
	if a < 0 || b < 0 {
		return errors.New("Apenas números positivos são aceitos")
	}

	resultado, err := dividir(a, b)

	if err != nil {
		return err
	}

	fmt.Println("Resultado da divisão: ", resultado)
	
	return nil
}

func limparLista() error {

	if len(numeros) == 0 {
		return errors.New("Lista está vazia")
	}

	numeros = []int{}
	fmt.Println("Lista limpa")	

	return nil
}

func listarNumerosPares() error {

	if len(numeros) == 0 {
		return errors.New("Lista está vazia")
	}

	for i := 0; i < len(numeros); i++ {
		if numeros[i]%2 == 0 {
			fmt.Println(numeros[i])
		} 
	}

	return nil
}


