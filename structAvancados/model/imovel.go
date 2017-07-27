package model

import "errors"

//Imovel para exemplo de utilização de ponteiros e structs
type Imovel struct {
	Y     int    `json:"cordenada_x"`
	X     int    `json:"cordenada_y"`
	Nome  string `json:"nome"`
	valor int    // campo não exportado então não pode atribuir etiqueta para json
}

var (
	//ErrValorImovelErrado Varoavel para retorno de erro
	ErrValorImovelErrado = errors.New("O valor informado é muito baixo")
	//ErrValorMuitoAlto Erro devido ao valor do imovél superior a media
	ErrValorMuitoAlto = errors.New("O valor informado é muito alto")
)

//Setvalor define o valor do imovel
func (i *Imovel) Setvalor(valor int) (err error) {
	i.valor = valor
	err = nil
	if i.valor < 50000 {
		err = ErrValorImovelErrado
		return
	} else if i.valor > 1500000 {
		err = ErrValorMuitoAlto
		return
	}
	return
}

//GetValor retorna o valor do imovel
func (i *Imovel) GetValor() int {
	return i.valor
}
