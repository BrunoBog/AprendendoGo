package main

func main() {
	messes := 6
	situacao := true
	cidade := "Teste"

	if messes <= 6 { // sem parêmteses e com os mesmos operadores gerais
		println("Este credor deve a pouco tempo ")
	}

	if cidade == "Teste" {
		println("Você ainda está testando né ?")
	}

	if situacao {
		println("Este credor está inadimplente")
	}

	if !situacao {
		println("Este credor está adimplente")
	} else {
		println("Este credor está inadimplente")
	}

	if desc, status := tempoDevedor(messes); status { // desta forma as variaveis não valem fora do escopo das {}
		println("O retorno da funcao tempo devedor é: ", desc, " e também: ", status)
	}

}

func tempoDevedor(messes int) (descricao string, status bool) {
	if messes > 0 {
		status = true
		descricao = "Está inadimplente"
		return
	}
	descricao = "O cliente está adimplente"
	return
}
