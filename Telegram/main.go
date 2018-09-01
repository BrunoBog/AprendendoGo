package main

import (
	"./model"
	"bufio"
	"encoding/json"
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"os"
)

func main() {
	config, err := carregarConfiguracoes()
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(config.KeyTelegram)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

}

func carregarConfiguracoes() (config model.Config, err error) {

	config = model.Config{}
	corpo, err := abreArquivo("Keys.json")
	if err != nil {
		return
	}

	dados := []byte(corpo)
	err = json.Unmarshal(dados, &config)
	if err != nil {
		println("Não consegui converter o Json de Configurações ", err.Error())
	}

	return
}

func abreArquivo(caminho string) (body string, err error) {

	arquivo, err := os.Open(caminho)
	if err != nil {
		fmt.Println("[abreArquivo] Erro ao abrir o arquivo ", err.Error())
	}
	defer arquivo.Close() // informa ao SO para fechar o arquivo apos o sistema termianar

	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		body += scanner.Text()
	}

	return
}
