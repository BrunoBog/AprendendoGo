package main

import (
	"fmt"
	"net/http"

	"github.com/AprendendoGo/Webservices/manipulador"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, " Ola mundo")
	})

	http.HandleFunc("/funcao", manipulador.FuncaoManipuladora)
	http.HandleFunc("/ola", manipulador.Ola)
	fmt.Println("O serviço subiu")
	http.ListenAndServe(":8181", nil)
}
