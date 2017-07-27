package model

//Galinha exemplo de uma interface
type Galinha interface {
	Carcareja() string
}

//Pata exemplo de uma interface
type Pata interface {
	Grasna() string
}

//Leao exemplo de uma interface
type Leao interface {
	Rugi() string
}

//Ave sao os tipos de aves
type Ave struct {
	Nome string
}

//Carcareja o que a ave faz ?
func (a Ave) Carcareja() string {
	return "cocorico"
}

//Grasna o que um pato faz
func (a Ave) Grasna() string {
	return "Quack"

}

//Rugi o que um leão faz
func (a Ave) Rugi() string {
	return "Rugido..."

}
