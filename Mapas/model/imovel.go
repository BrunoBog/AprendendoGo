package model

//Imovel para exemplo de utilização de ponteiros e structs
type Imovel struct {
	Y     int    `json:"cordenada_x"`
	X     int    `json:"cordenada_y"`
	Nome  string `json:"nome"`
	valor int    // campo não exportado então não pode atribuir etiqueta para json
}

//Setvalor define o valor do imovel
func (i *Imovel) Setvalor(valor int) {
	i.valor = valor
}

//GetValor retorna o valor do imovel
func (i *Imovel) GetValor() int {
	return i.valor
}
