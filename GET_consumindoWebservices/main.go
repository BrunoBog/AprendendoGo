package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}
	//get simples
	resposta, err := cliente.Get("https://www.google.com.br")
	if err != nil {
		println("[main] Erro ao abrir a URL\n", err.Error())
		return
	}

	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		println("consegui obter uma resposta")
		println(resposta.Status)

		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			println("[main] Erro ao ler o corpo da resposta")
			return
		}

		println(corpo)
	}
	// estruturando o request com autenticação
	request, err := http.NewRequest("GET", "https://www.google.com.br", nil)
	if err != nil {
		println("[main], deu erro ao tentar fazer o request na pagina do google")
	}
	//usuario e senha aqui
	request.SetBasicAuth("teste", "teste")
	resposta, err = cliente.Do(request)
	if err != nil {
		println("[main], deu erro ao tentar executar o request")
	}

	if resposta.StatusCode == 200 {
		println("consegui obter uma resposta")
		println(resposta.Status)

		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			println("[main] Erro ao ler o corpo da resposta")
			return
		}

		println(string(corpo))
	}
	defer resposta.Body.Close()

}
