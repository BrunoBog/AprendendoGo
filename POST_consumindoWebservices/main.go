package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/AprendendoGo/POST_consumindoWebservices/model"
)

func main() {
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}
	usuario := model.Usuario{}
	usuario.ID = 1
	usuario.Nome = "Bruno"

	usuarioJSON, err := json.Marshal(usuario)
	if err != nil {
		println("[main] Rolou um erro ao converter usuario para json")
	}

	// estruturando o request com autenticação
	request, err := http.NewRequest("POST", "https://requestb.in/11rysp61", bytes.NewBuffer(usuarioJSON))
	if err != nil {
		println("[main], deu erro ao tentar fazer o request para o requestbin")
	}
	//usuario e senha aqui
	request.SetBasicAuth("fizz", "buzz")
	// para enviar um content type temos que add ele noo request
	request.Header.Set("content-type", "application/json; charset=utf-8")
	resposta, err := cliente.Do(request)
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
