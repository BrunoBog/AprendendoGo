package model

//BlogPost armazena dados de um post emum blog
type BlogPost struct {
	UsuarioID int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
}
