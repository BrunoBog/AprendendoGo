package main

import "github.com/AprendendoGo/structAvancados/model"
import "fmt"
import "encoding/json"

func main() {

	casa := model.Imovel{}
	casa.X = 50
	casa.Y = 60
	casa.Nome = "Minha casa"
	if err := casa.Setvalor(12500000000); err != nil {
		fmt.Println("[main] Houve um erro na atribuição do valor. \n ", err.Error())
		if err == model.ErrValorMuitoAlto {
			fmt.Println("[main]Por favor confira o valor digitado ou sua avaliação")
		}

		return
	}

	fmt.Println("O valor de casa é: ", casa)
	fmt.Printf("O valor da casa é: %d\n", casa.GetValor)   // assim você chama o endereço de memoria
	fmt.Printf("O valor da casa é: %d\n", casa.GetValor()) // assim você chama o valor
	//objJson, _ := json.Marshal(casa)                       // transforma a parada em json
	// você usa o "_" ao invés do err quando quer ignorar o retorno do erro por exemplo
	objJSON, err := json.Marshal(casa)
	if err != nil {
		fmt.Println("[main] Deu erro na hora de gerar o jSon T_T", err.Error)
		return //panic se retornar panic o app fecha
	}
	fmt.Println("O valor da casa em json", string(objJSON))

}
