package main

import (
	"fmt"
	"net/http"

	"Webservices/manipulador"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, " Ola mundo")
		//Testar pra ver qual retorna valor
		// param1 := r.URL.Query().Get("parametro1")
		param2 := r.FormValue("password")
		if param2 != nil {
			fmt.Println("Tinha parametro")
		} else {
			fmt.Println("Não achei parametro")
		}
	})

	http.HandleFunc("/funcao", manipulador.FuncaoManipuladora)
	http.HandleFunc("/ola", manipulador.Ola)
	fmt.Println("O serviço subiu")
	http.ListenAndServe(":8181", nil)

}
