package operadora

import "github.com/AprendendoGo/pacotes/prefixo"
import "strconv"

//Operadora variavél para guardar a operadora
var operadora = "Oi "

//PrefixoOperadora variavél para guardar a operadora e o DDD
var PrefixoOperadora = operadora + " " + strconv.Itoa(prefixo.Capital)

//str conv serve para converter os dados
