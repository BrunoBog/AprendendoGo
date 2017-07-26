package main

import "fmt"

//Imovel estrutura para testes
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {
	casa := Imovel{}
	fmt.Printf("A casa é: %+v\n", casa)

	apartamento := Imovel{1000, 1000, " Meu Ap", 760000}
	fmt.Printf("O apartamenteo é: %+v\n", apartamento)

	chacara := Imovel{
		X:     2500,
		Y:     2500,
		Nome:  " O sitio",
		valor: 1005000,
	}

	fmt.Printf("O apartamenteo é: %+v\n", chacara)

	casa.Nome = "minha casa"
	casa.valor = 350000
	casa.X = 50
	casa.Y = 60

	fmt.Printf(" Os novos valores da casa são: %+v", casa)
}
