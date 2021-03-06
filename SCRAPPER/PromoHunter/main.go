package main

import (
	"./model"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	orquestrador sync.WaitGroup
)

func HardMobScrap() (Promocoes model.HardMobPromo) {

	doc, err := goquery.NewDocument("http://www.hardmob.com.br/promocoes/")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div .inner").Each(func(i int, s *goquery.Selection) {
		// For each item...
		item := model.HardMopbItem{}
		item.Nome = strings.Replace((strings.Replace(s.Find("a").Text(), "\n", "", -1)), "\t", "", -1)
		link, _ := s.Find("a").Attr("href")
		item.Link = strings.Replace(link, " ", "", -1)
		index := strings.Index(item.Nome, "R$")
		if index > 0 {
			item.Preco = item.Nome[index:(len(item.Nome) - 1)]
		}
		Promocoes.Oferta = append(Promocoes.Oferta, item)

	})
	return
}

func compareAndSendHardMob(desejo string, item model.HardMopbItem, caminho string) {
	if strings.Contains(strings.ToUpper(item.Nome), strings.ToUpper(desejo)) {
		fmt.Println(item.Nome, item.Link)
		json, _ := json.Marshal(item)
		orquestrador.Add(1)
		go doRequest(json, caminho)
	}
	orquestrador.Done()
}

func filtrarDesejoHardMob(desejo string, caminho string) (err error) {
	for _, item := range HardMobScrap().Oferta {
		fmt.Println(item.Nome, item.Link)
		orquestrador.Add(1)
		go compareAndSendHardMob(desejo, item, caminho)
	}
	orquestrador.Done()
	return
}

func PromoBitScrap() (Promocoes model.PromobitPromo) {

	doc, err := goquery.NewDocument("https://www.promobit.com.br/")
	if err != nil {
		log.Fatal(err)
	}
	// digite a tag e .nome da class
	doc.Find("div .pr-tl-card").Each(func(i int, s *goquery.Selection) {
		// For each item...
		item := model.PromobitItem{}
		item.Nome = strings.Replace((strings.Replace(s.Find("a").Text(), "\n", "", -1)), "\t", "", -1)
		item.Preco = s.Find("div .price").Text()
		item.Loja = s.Find("div .where").Text()
		link, _ := s.Find("a").Attr("href")
		item.Link = "www.promobit.com.br" + strings.Replace(link, " ", "", -1)

		Promocoes.Oferta = append(Promocoes.Oferta, item)

	})
	return
}

func compareAndSendPromobit(desejo string, item model.PromobitItem, caminho string) {
	log.Println(item.Nome, item.Link)
	if strings.Contains(strings.ToUpper(item.Nome), strings.ToUpper(desejo)) {
		fmt.Println(item.Nome, item.Link)
		json, _ := json.Marshal(item)
		orquestrador.Add(1)
		go doRequest(json, caminho)
	}
	orquestrador.Done()
}

func filtrarDesejoPromobit(desejo string, caminho string) (err error) {
	for _, item := range PromoBitScrap().Oferta {
		orquestrador.Add(1)
		go compareAndSendPromobit(desejo, item, caminho)
	}
	orquestrador.Done()
	return
}

func doRequest(json []byte, caminho string) {

	cliente := &http.Client{
		Timeout: time.Second * 30,
	}
	// estruturando o request com autenticação
	request, err := http.NewRequest("POST", caminho, bytes.NewBuffer(json))
	if err != nil {
		return
		//println("[doRequest], erro ao tentar fazer o request para o requestbin")
	}
	//usuario e senha aqui
	request.SetBasicAuth("fizz", "buzz")
	// para enviar um content type temos que add ele noo request
	request.Header.Set("content-type", "application/json; charset=utf-8")
	resposta, err := cliente.Do(request)
	defer resposta.Body.Close()
	if err != nil {
		orquestrador.Done()
		return
		//println("[doRequest], deu erro ao tentar executar o request")
	}
	if resposta.StatusCode == 200 {
		//println(resposta.Status)

		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			orquestrador.Done()
			//println("[doRequest] Erro ao ler o corpo da resposta")
			return
		}
		fmt.Println(string(corpo))
	}

	orquestrador.Done()
}

//addDesejo  -pendente
func addDesejo(json []byte) {

	//desejo := model.Desejos{}
	//json.Unmarshal(json, &desejo)

	//desejo.AddDesejos()

	return
}

func main() {

	orquestrador.Add(1)
	go filtrarDesejoHardMob("Luva", "http://requestbin.fullcontact.com/zp13d1zp") // item - destino
	//go filtrarDesejoPromobit("Luva", "http://requestbin.fullcontact.com/zp13d1zp") // item - destino
	orquestrador.Wait()

}
