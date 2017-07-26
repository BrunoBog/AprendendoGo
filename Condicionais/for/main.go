package main

import (
	"fmt"
)

func main() { // só existe o loop do for kkkk
	for i := 0; i <= 10; i++ {
		fmt.Println("o valor de i é: ", i)
	}

	valor := 0
	teste := true

	for teste {
		valor++
		if valor%5 == 0 {
			teste = false
		}
		fmt.Println("O valor de valor é: ", valor, " e o valor de teste é: ", teste)
	}

	for {
		valor--
		fmt.Println("O valor de valor é: ", valor)
		if valor == 0 {
			break
		}
	}

	text := "eu estou escrevendo muita coisa pra informar quse nada T_T"
	for indice, letra := range text {
		// se eu não usar string() ele retorna asc
		fmt.Println("O indice do texto é: ", indice, " e a letra é: ", string(letra))
	}
	
}
