package model

import (
	"database/sql"
)

//Local armazena os dados da localização pelo codigo de area (DDD)
type Local struct {
	Pais             string         `json:"pais" db:"pais"`
	Cidade           sql.NullString `json:"cidade" db:"cidade"`
	CodigoTelefonico int            `json:"cod_fone" db:"cod_fone"`
}
