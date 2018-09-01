package main

import (
	"errors"
	"fmt"
)

var errDivisaoInvalida = errors.New("divisão invalida")

func main() {
	fmt.Println("subiu")
}

func divideInteiros(dividendo, divisor int) (quociente int, resto int, err error) {
	if divisor == 0 {
		err = errDivisaoInvalida
		return
	}
	quociente = dividendo / divisor
	resto = dividendo % divisor
	return
}
