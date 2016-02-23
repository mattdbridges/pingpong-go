package main

import (
	"fmt"
)

func ponger(c chan string) {
	for {
		select {
		case msg := <-c:
			switch msg {
			case "ping":
				fmt.Println("pong")
			case "pong":
				fmt.Println("ping")
			}
		}
	}
}

func asker(pinger chan<- string, done chan<- bool) {
	var input string
	fmt.Scanln(&input)
	switch input {
	case "ping", "pong":
		pinger <- input
		asker(pinger, done)
	case "exit":
		done <- true
	}
}

func main() {
	done := make(chan bool, 1)
	channel := make(chan string, 1)

	go ponger(channel)
	go asker(channel, done)

	<-done
}
