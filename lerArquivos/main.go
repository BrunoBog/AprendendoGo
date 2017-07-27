package main

import "os"
import (
	"fmt"
	"bufio"
)

func main() {
	arquivo, err := os.OpenFile("cidades.csv")
	if err != nil {
		fmt.Println("[main] Erro ao abrir o arquivo ", err.Error())
	}

	fmt.Printf("Arquivo contém: ",arquivo)
	scanner :=  bufio.NewScanner(arquivo)
	for scanner.Scan(){
	linha := scanner.Text()
		fmt.Printf("linha contém: ",linha)
	}
}
