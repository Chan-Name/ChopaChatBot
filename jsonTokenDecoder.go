package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Token struct {
	Token string
}

func jsonTokenDecode() string {
	file, err := os.Open("token.json")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var token Token

	err = json.Unmarshal(bytes, &token)
	if err != nil {
		log.Fatal(err)
	}

	return token.Token
}
