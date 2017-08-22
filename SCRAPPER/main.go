package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func meuscrap() {
	doc, err := goquery.NewDocument("http://carteirarica.com.br/taxa-selic/")
	if err != nil {
		log.Fatal(err)
	}
	title := ""
	// digite a tag e .nome da class
	doc.Find("ul .price").Each(func(i int, s *goquery.Selection) {
		// For each item...
		title = s.Find("span").Text()
	})
	fmt.Println(title)
}

func main() {
	meuscrap()
}
