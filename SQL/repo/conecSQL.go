package repo

import (
	"github.com/jmoiron/sqlx"
	//_ "github.com/mattn/go-sqlite3"
	// usa o _ para dizer que queremos que as variaveis de init e variaveis seráo inicializadas
	_ "github.com/go-sql-driver/mysql"
)

var Db *sqlx.DB
