package main

import (
	"fmt"

	"github.com/AprendendoGo/structAvancados/model"
)

var cache map[string]model.Imovel // quando você cria um mapa ele fica como Null

func main() {

	casa := model.Imovel{}
	casa.Nome = "Cazão"
	casa.X = 18
	casa.Y = 25
	casa.Setvalor(350000)

	cache = make(map[string]model.Imovel, 0) // isso aqui é pra inicializar o mapa

	cache["Casa Grande"] = casa
	casaPequena := model.Imovel{Nome: "Casa Pequena", X: 10, Y: 2}
	casaPequena.Setvalor(200000)
	cache["casa Pequena"] = casaPequena

	fmt.Println("Verificando se há uma casa Grande...")
	imovel, achou := cache["Casa Grande"]
	if achou {
		fmt.Println("Encontrei este imovel:", imovel)
	} else {
		fmt.Println("Tem não viu ...")
	}

	println("O tamanho do cache é: ", len(cache))

	//fazer buscas....
	for chave, imovel := range cache {
		println("A chave ", chave, " é igual a: ", imovel.Nome)
	}

	// para remover item...
	_, achou = cache["Casa Grande"] // se usa o _ para ignorar o retorno
	if achou {
		delete(cache, imovel.Nome)
		//delete(cache,"Casa Grande")
	}
	println("Ainda restam ", len(cache), " Itens no cache")

}
