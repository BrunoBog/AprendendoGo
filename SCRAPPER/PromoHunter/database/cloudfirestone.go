package main

//https://firebase.google.com/docs/firestore/quickstart?hl=pt-br

import (
	"log"

	firebase "firebase.google.com/go"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {

	// Use a service account
	sa := option.WithCredentialsFile("chave.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	gravarDados(client)
	lerDados(client)

}

// I Don´t know how do this.... https://godoc.org/firebase.google.com/go#App
func gravarDados(client *firestore.Client) {
	/* Para adicionar dados */
	_, _, err := client.Collection("users").Add(context.Background(), map[string]interface{}{
		"first": "Ada",
		"last":  "Lovelace",
		"born":  1815,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

	_, _, err = client.Collection("users").Add(context.Background(), map[string]interface{}{
		"first":  "Alan",
		"middle": "Mathison",
		"last":   "Turing",
		"born":   1912,
	})
	if err != nil {
		log.Fatalf("Failed adding aturing: %v", err)
	}

}

func lerDados(client *firestore.Client) {
	iter := client.Collection("users").Documents(context.Background())

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		log.Println(doc.Data())
	}
}
