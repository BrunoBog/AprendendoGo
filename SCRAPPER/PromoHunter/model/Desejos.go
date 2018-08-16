package model

//Desejos representa os desejos de um Usuario
type Desejos struct {
	Email   string   `json:"Email" bson:"Email"`
	Desejos []Desejo `json:"Desejos" bson:"Desejos"`
}

//Desejo Representa os itens que o usuario quer comrar
type Desejo struct {
	Nome string `json:"Nome" bson:"Nome"`
}

func (d *Desejos) BuscarDesejos(email string) (err error) {

	return
}

func (d *Desejos) AddDesejos(email string) (err error) {

	return
}
