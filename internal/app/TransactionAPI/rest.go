package TransactionAPI

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var DBServer DBQuerier

func (t Transaction) createTransaction(config Config, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requested create transaction")
	DBServer = t
	err := DBServer.CreateTransaction(config)
	if err != nil {
		log.Println(err)
	}
	b, err := fmt.Fprint(w, err)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(b)
}

func (t Transaction) getStatusById(config Config, w http.ResponseWriter, r *http.Request) {
	splitRoute := strings.Split(r.URL.String(), "?id=")
	id, err := strconv.Atoi(splitRoute[1])
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Requested status of transaction id: %d\n", id)
	DBServer = t
	status, err := DBServer.GetTransactionStatusById(config, id)
	if err != nil {
		log.Println(err)
	}
	b, err := fmt.Fprint(w, status)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(b)
}

func (t Transaction) getAllTransactionsByUserId(config Config, w http.ResponseWriter, r *http.Request) {
	splitRoute := strings.Split(r.URL.String(), "?id=")
	id, err := strconv.Atoi(splitRoute[1])
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Requested all transactions of UserId: %d\n", id)
	DBServer = t
	allTransactionsByUserId, err := DBServer.GetAllTransactionsByUserId(config, id)
	if err != nil {
		log.Println(err)
	}
	b, err := fmt.Fprint(w, allTransactionsByUserId)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(b)
}

func (t Transaction) getAllTransactionsByUserEmail(config Config, w http.ResponseWriter, r *http.Request) {
	splitRoute := strings.Split(r.URL.String(), "?email=")
	email := strings.Trim(splitRoute[1], "%22")
	fmt.Printf("Requested all transactions of UserEmail: %v\n", email)
	DBServer = t
	allTransactionsByUserEmail, err := DBServer.GetAllTransactionsByUserEmail(config, email)
	if err != nil {
		log.Println(err)
	}
	b, err := fmt.Fprint(w, allTransactionsByUserEmail)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(b)
}
