package main

import (
	"fmt"
	"bytes"
	"encoding/json"
	"log"
)

type album struct {
	ID		string	`json:"id"`
	Title	string	`json:"title"`
	Artist	string	`json:"artist"`
	Price	float64	`json:"price"`
}

func main() {
	a := album{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99}
	b, err := json.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}

	b2 := []byte(`{"id":"1","title":"Blue Train","artist":"John Coltrane","price":56.99}`)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(bytes.Equal(b, b2))

	var a2 album
	err = json.Unmarshal(b2, &a2) // b2 []byte -> a2 album
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a2)
}
