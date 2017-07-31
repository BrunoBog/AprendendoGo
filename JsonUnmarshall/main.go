package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/AprendendoGo/JsonUnmarshall/model"
)

func main() {
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}
	//url fixa do site
	modificador := "/posts/1"
	// estruturando o request com autenticação
	request, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com"+modificador, nil)
	if err != nil {
		println("[main], deu erro ao tentar fazer o request na pagina do servidor" + err.Error())
	}
	//usuario e senha aqui
	request.SetBasicAuth("teste", "teste")
	resposta, err := cliente.Do(request)
	if err != nil {
		println("[main], deu erro ao tentar executar o request" + err.Error())
	}

	if resposta.StatusCode == 200 {
		println("consegui obter uma resposta")
		println(resposta.Status)

		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			println("[main] Erro ao ler o corpo da resposta" + err.Error())
			return
		}

		println(" ")
		post := model.BlogPost{}
		err = json.Unmarshal(corpo, &post)
		if err != nil {
			println("Não consegui converter o Json ", err.Error())
		}

		fmt.Println("O resultado do Json é: ", post)
	}
	defer resposta.Body.Close()

}
