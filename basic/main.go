package main

import (
	"fmt"
	"log"

	"github.com/devtae/go-exercise/greetings"
)

func main() {
	message, err := greetings.Hello("Taehyeon")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}