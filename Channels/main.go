package main

import (
	"fmt"
	"time"
)

func pong(canal chan string) {
	for {
		canal <- "ping"
	}
}

func ping(canal chan string) {
	for {
		canal <- "pong"
	}
}

func impressora(canal chan string) {
	for {
		msg := <-canal
		println(msg)
		time.Sleep(time.Second * 2)
	}

}
func main() {
	var canal chan string
	canal = make(chan string)

	go ping(canal)
	go pong(canal)
	go impressora(canal)

	var entrada string
	fmt.Scanln(&entrada)

}
