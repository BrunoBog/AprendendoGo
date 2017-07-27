package main

import (
	"fmt"
)

func main() {
	var teste [3]int
	teste[0] = 3
	teste[0] = 2
	teste[0] = 1

	fmt.Println("O vetor contém: ", len(teste), " posições")

	cantores := [2]string{"michael jackson", "Bruce Dickinson "}
	fmt.Println("O vetor cantores contém: ", cantores)

	capitais := [...]string{"Brasilia", "luanda", "Maputo", "lisboa"}
	fmt.Println("Capitais que falam portugês: \n", capitais)

	for indice, cidade := range capitais {
		fmt.Println("Capital[", indice, "]: ", cidade)
	}

}
