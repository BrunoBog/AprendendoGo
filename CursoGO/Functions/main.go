package main

import (
	"strconv"

	"github.com/AprendendoGo/CursoGO/Functions/matematica"
)

func main() {
	x := 2
	y := 2
	resultado := multiplicador(2, 2)
	println("Resultado de " + strconv.Itoa(x) + " + " + strconv.Itoa(y) + " = " + strconv.Itoa(resultado))
	novoResultado := matematica.Soma(x, y)
	println("resultado da function " + strconv.Itoa(novoResultado))
	resultado3 := matematica.Calculo(matematica.Multiplicador, 2, 2)
	println("resultado da função Calculo: multiplicando x e y" + strconv.Itoa(resultado3))
}

func multiplicador(x int, y int) int {
	return x * y
}
