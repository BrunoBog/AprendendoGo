package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	numero := 3
	fmt.Println("O numero", numero, "se escreve assim:")

	switch numero {
	case 1:
		fmt.Println("Um")
	case 2:
		fmt.Println("dois")
	case 3:
		fmt.Println("Tres")
	}

	fmt.Println("Essa parada só compila em utf-8 \\o/")
	switch runtime.GOOS { // checando a versão do sistema
	case "darwin":
		fallthrough // isso aqui é tipo um break, que você vai pro proximo
	case "freeBSD":
		fallthrough
	case "windows":
		fmt.Println("Sou windows")
	case "linux":
		fmt.Println("Sou linux")
	default:
		fmt.Println("Sou a oppção padrão....")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Sabado ou Quarta")
	default:
		fmt.Println("Dia de trabalhar cara!!")
	}

	switch { // entra na 1° condição e se for satisfeito ele sai do switch
	case numero > 2:
		fmt.Println("O numero é maior que 2")

	case numero > 1, numero < 100:
		fmt.Println("O numero é maior que 1 e menor que 100")
	}
}
