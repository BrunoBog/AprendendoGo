package main

import (
	"fmt"
)

var (
	/* se a variavél começar com  letra maiuscula quer dizer que ela é publica
	mas se começar com letra minuscula é private*/
	endereco string // vazia é = ''
	// Telefone é um valor global do n° quando fizer uma variavel publica o GO pede pra add um comentario... O comentario deve começar com o nome da var, Ex:
	Telefone   = "31 3561-2485" // exemplo de var Publica
	quantidade int              // vazio é igual a 0
	comprou    bool             // vazio é false
	valor      float64          // vazio 0.00
	palavras   rune
)

// se você declarar uma variavel e ela não estiver sendo usada o compulador não compila o código
func main() {
	fmt.Printf("endereço %s\n", endereco)
	fmt.Printf("Quantidade %d\n", quantidade)
	fmt.Printf("comprou %v\n", comprou)
	fmt.Printf("Runas: %v\n", palavras)
	teste := "valor do teste"
	fmt.Println("o teste é: " + teste)
}
