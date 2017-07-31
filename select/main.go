package main

// select é um recurso para capturar o que há no canal entre as slices
import (
	"fmt"
	"time"
)

var (
	irc = make(chan string)
	sms = make(chan string)
)

func pong(canal chan string) {
	for {
		canal <- "ping"
	}
}

func eaew(canal chan string) {
	for {
		canal <- "e aew"
	}
}

func opa(canal chan string) {
	for {
		canal <- "WAZAAAAP"
	}
}

func ping(canal chan string) {
	for {
		canal <- "pong"
	}
}

func impressora() {
	var msg string
	for {
		select {
		case msg = <-irc:
			println("Mensagem por irc ", msg, "old but gold")
			time.Sleep(time.Second * 2)

		case msg = <-sms:
			println("Mensagem por SMS ", msg, "ZZZzzzZZZ \n")
			time.Sleep(time.Second * 2)
		}

	}

}
func main() {

	go ping(irc)
	go pong(irc)
	go eaew(sms)
	go opa(sms)
	go impressora()

	var entrada string
	fmt.Scanln(&entrada)

}
