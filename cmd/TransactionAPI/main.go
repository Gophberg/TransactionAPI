package main

import (
	"github.com/Gophberg/TransactionAPI/internal/app/TransactionAPI"
	"log"
)

func main() {
	if err := TransactionAPI.Start(); err != nil {
		log.Fatal(err)
	}
}
