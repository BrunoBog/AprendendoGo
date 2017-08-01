package manipulador

import (
	"fmt"
	"net/http"
	"time"

	"github.com/AprendendoGo/Webservices/model"
)

//Ola é o Manipulador da requisição que ira gerar o html da rota ola
func Ola(w http.ResponseWriter, r *http.Request) {
	hora := time.Now().Format("15:15:05")
	pagina := model.Pagina{}
	pagina.Hora = hora
	if err := Modelos.ExecuteTemplate(w, "ola.html", pagina); err != nil {
		http.Error(w, "Houve um erro na renderização da pagina", http.StatusInternalServerError)
		fmt.Println("[ola] Deu erro na bagaça toda velho.... \n", err.Error())
	}
}
