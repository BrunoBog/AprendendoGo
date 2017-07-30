package main

import "os"
import (
	"fmt"
	"bufio"
)

func main() {
	arquivo, err := os.Open("cidades.csv")
	if err != nil {
		fmt.Println("[main] Erro ao abrir o arquivo ", err.Error())
	}
	defer arquivo.Close() // informa ao SO para fechar o arquivo apos o sistema termianar

	fmt.Printf("Arquivo contém: ",arquivo)
	scanner :=  bufio.NewScanner(arquivo)
	for scanner.Scan(){
	linha := scanner.Text()
		fmt.Printf("linha contém: ",linha)
	}

	arquivo.Close() // não é necessario pois já usou o defer antes
}
