package model

//Usuario Representa um user na estrutura
type Usuario struct {
	//ID do usuario
	ID   int    `json:"ID"`
	Nome string `Json:"nome"`
}
