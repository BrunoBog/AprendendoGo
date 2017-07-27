package main

import "github.com/AprendendoGo/interfaces/model"
import "fmt"

func main() {
	jojo := model.Ave{}
	jojo.Nome = "Agnes"

	acordandoComGalinha(jojo)
	patosBerrando(jojo)
	rugir(jojo)

}

func acordandoComGalinha(g model.Galinha) {
	fmt.Println(g.Carcareja())
}

func patosBerrando(p model.Pata) {
	fmt.Println(p.Grasna())
}

func rugir(l model.Leao) {
	fmt.Println(l.Rugi())
}
