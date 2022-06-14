package main

import (
	"TransactionAPI/internal/app/TransactionAPI"
	"log"
)

func main() {
	config := TransactionAPI.NewConfig()
	if err := TransactionAPI.Start(config); err != nil {
		log.Fatal(err)
	}
}
