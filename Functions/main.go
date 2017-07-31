package main

import (
	"strconv"

	"github.com/AprendendoGo/Functions/matematica"
)

func main() {

	x := 2
	y := 2
	resultado := multiplicador(2, 2)
	println("Resultado de " + strconv.Itoa(x) + " + " + strconv.Itoa(y) + " = " + strconv.Itoa(resultado))
	novoResultado := matematica.Soma(x, y)
	println("resultado da function " + strconv.Itoa(novoResultado))
	resultado = matematica.Calculo(matematica.Multiplicador, 2, 2) // como já foi inicializado não usa os :
	println("resultado da função Calculo: multiplicando x e y " + strconv.Itoa(resultado))
	resultado, resto := matematica.Divisor(2, 5)
	println("O resultado da divisão é: ", resultado, " e o resto: ", resto)
}

func multiplicador(x int, y int) int {
	return x * y
}
