package main

import (
	"fmt"
)

//Imovel estrutura para testes
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {
	casa := new(Imovel) // assim você está pegando o endereço de
	fmt.Println("O ponteiro de Imovel tem o valor de: ", casa)
	fmt.Println("e o endereço do ponteiro é:", &casa)

	chacara := Imovel{
		17, 28, "chacara", 280000,
	}
	fmt.Println("\nO ponteiro de chacara tem o valor de: ", chacara)
	fmt.Println("e o endereço do ponteiro de chacara é:", &chacara)

	mudaImovel(&chacara)
}

// o Go sempre passa aas variaveis como valor se quiser usar referencia tem q usar o * para indicar que é um vetor
func mudaImovel(imovel *Imovel) {
	imovel.X = imovel.X + 10
	imovel.Y += 10
}
