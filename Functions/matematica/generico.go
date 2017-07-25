package matematica

//Calculo é para executar qualquer tipo de calculo, basta enviar a função
func Calculo(funcao func(int, int) int, numeroA int, numeroB int ) int{
	return funcao(numeroA,numeroB)
}

//Multiplicador Apenas multiplica os 2 argumentos
func  Multiplicador(x int, y int)int  {
	return x * y
}

//Divisor Apenas divide os 2 argumentos e devolve o resultado e o resto da divisão
func Divisor(x int, y int) (resultado int, resto int)  {
	resultado = x/y
	resto = x%y
	return
}

