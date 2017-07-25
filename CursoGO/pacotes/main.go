package main

import (
	"fmt"
	//"github.com/brunobog/CursoGO/pacotes/prefixo"
	"github.com/AprendendoGo/CursoGO/pacotes/operadora"
	"github.com/AprendendoGo/CursoGO/pacotes/prefixo"
)

//NomeDoUsuario Nome do usuario do sistema
var NomeDoUsuario = "BOG"

func main() {
	fmt.Printf("Nome do user: " + NomeDoUsuario + "\n")
	fmt.Println("Prefixo da capital: ", prefixo.Capital)
	fmt.Println("Variavel concatenada : " + operadora.PrefixoOperadora)
}
