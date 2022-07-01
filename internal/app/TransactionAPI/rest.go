package TransactionAPI

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (t Transaction) createTransaction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requested create transaction")
	//fmt.Printf("Request: %v\n", r)
	err := t.CreateTransaction()
	if err != nil {
		log.Println(err)
	}
	b, err := fmt.Fprint(w, err)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(b)
}

func (t Transaction) getTransactionStatusById(w http.ResponseWriter, r *http.Request) {
	splitRoute := strings.Split(r.URL.String(), "?id=")
	id, err := strconv.Atoi(splitRoute[1])
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Requested status of transaction id: %d\n", id)

	status, err := t.GetTransactionStatusById(id)
	if err != nil {
		log.Println(err)
	}
	b, err := fmt.Fprint(w, status)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(b)
}

func (t Transaction) getAllTransactionsByUserId(w http.ResponseWriter, r *http.Request) {
	splitRoute := strings.Split(r.URL.String(), "?id=")
	id, err := strconv.Atoi(splitRoute[1])
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Requested all transactions of UserId: %d\n", id)
	transactions, err := t.GetAllTransactionsByUserId(id)
	if err != nil {
		log.Println(err)
	}
	b, err := fmt.Fprint(w, transactions)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(b)
}

func (t Transaction) getAllTransactionsByUserEmail(w http.ResponseWriter, r *http.Request) {
	splitRoute := strings.Split(r.URL.String(), "?email=")
	email := strings.Trim(splitRoute[1], "%22")
	fmt.Printf("Requested all transactions of UserEmail: %v\n", email)
	transactions, err := t.GetAllTransactionsByUserEmail(email)
	if err != nil {
		log.Println(err)
	}
	b, err := fmt.Fprint(w, transactions)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(b)
}
