package main

import (
	"fmt"
	"github.com/AprendendoGo/pacotes/operadora"
	"github.com/AprendendoGo/pacotes/prefixo"
)

//NomeDoUsuario Nome do usuario do sistema
var NomeDoUsuario = "BOG"

func main() {
	fmt.Printf("Nome do user: " + NomeDoUsuario + "\n")
	fmt.Println("Prefixo da capital: ", prefixo.Capital)
	fmt.Println("Variavel concatenada : " + operadora.PrefixoOperadora)
}
