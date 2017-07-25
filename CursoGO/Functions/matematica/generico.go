package matematica

//Calculo é para executar qualquer tipo de calculo, basta enviar a função
func Calculo(funcao func(int, int) int, numeroA int, numeroB int ) int{
	return funcao(numeroA,numeroB)
}

//Multiplicador Apenas multiplica os 2 argumentos
func  Multiplicador(x int, y int)int  {
	return x * y
}

