<h1># AprendendoGo</h1>
<hr></hr>
<h3>Implementações simples para exercitar o aprendizado</h3>

<h4>Para Gerar o build:</h4>  
go build -o ("nome do app")   
<h4>Para gerar cross plataform: <h4>
<h5>Windows</h5>
GOOS=windows GOARCH=386 go build -o ("nome do app")  
<h5>Linux</h5>
GOOS=linux GOARCH=arm go build -o ("nome do app")  -v //para lançar um verboose



