package main

import "github.com/AprendendoGo/structAvancados/model"
import "fmt"
import "encoding/json"

func main() {

	casa := model.Imovel{}
	casa.X = 50
	casa.Y = 60
	casa.Nome = "Minha casa"
	casa.Setvalor(75000)

	fmt.Println("O valor de casa é: ", casa)
	fmt.Printf("O valor da casa é: %d\n", casa.GetValor)   // assim você chama o endereço de memoria
	fmt.Printf("O valor da casa é: %d\n", casa.GetValor()) // assim você chama o valor
	objJson, _ := json.Marshal(casa)                       // transforma a parada em json
	// você usa o "_" quando quer ignorar um retorno
	fmt.Println("O valor da casa em json", string(objJson))
}
