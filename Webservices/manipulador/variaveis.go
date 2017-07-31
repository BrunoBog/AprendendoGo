package manipulador

import "html/template"

//Modelos armazena o modelo de pagina que será executado pelis manipuladorres
var Modelos = template.Must(template.ParseFiles("html/ola.html"))
