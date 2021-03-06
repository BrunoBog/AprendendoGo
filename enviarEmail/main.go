package main

import (
	"log"
	"net/smtp"
)

func main() {
	send("hello there")
}

func send(body string) {
	from := "procuromuamba@gmail.com"
	pass := "sua senha"
	to := "teste@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
}
