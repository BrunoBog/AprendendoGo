package manipulador

import (
	"fmt"
	"net/http"
)

//FuncaoManipuladora é o manipulador da rota / Função
func FuncaoManipuladora(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Aqui é o manipulador de dentro do pacote")
}
