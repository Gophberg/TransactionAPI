package main

import (
	"TransactionAPI/internal/app/TransactionAPI"
	"log"
)

func main() {
	if err := TransactionAPI.Start(); err != nil {
		log.Fatal(err)
	}
}
