package TransactionAPI

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var DBServer DBQuerier

func (t Transaction) createTransaction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requested create transaction")
	DBServer = t
	err := DBServer.CreateTransaction()
	if err != nil {
		log.Println(err)
	}
	b, err := fmt.Fprint(w, err)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(b)
}

func (t Transaction) getStatusById(w http.ResponseWriter, r *http.Request) {
	splitRoute := strings.Split(r.URL.String(), "?id=")
	id, err := strconv.Atoi(splitRoute[1])
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Requested status of transaction id: %d\n", id)
	DBServer = t
	status, err := DBServer.GetTransactionStatusById(id)
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
	DBServer = t
	allTransactionsByUserId, err := DBServer.GetAllTransactionsByUserId(id)
	if err != nil {
		log.Println(err)
	}
	b, err := fmt.Fprint(w, allTransactionsByUserId)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(b)
}

func (t Transaction) getAllTransactionsByUserEmail(w http.ResponseWriter, r *http.Request) {
	splitRoute := strings.Split(r.URL.String(), "?email=")
	email := strings.Trim(splitRoute[1], "%22")
	fmt.Printf("Requested all transactions of UserEmail: %v\n", email)
	DBServer = t
	allTransactionsByUserEmail, err := DBServer.GetAllTransactionsByUserEmail(email)
	if err != nil {
		log.Println(err)
	}
	b, err := fmt.Fprint(w, allTransactionsByUserEmail)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(b)
}

func Start(config *Config) error {
	var rustServer RestServer
	var DBServer = Transaction{}
	rustServer = DBServer
	http.HandleFunc("/getStatus", rustServer.getStatusById)
	http.HandleFunc("/getAllTransactionsByUserId", rustServer.getAllTransactionsByUserId)
	http.HandleFunc("/getAllTransactionsByUserEmail", rustServer.getAllTransactionsByUserEmail)
	//t := Transaction{
	//	UserID:       5,
	//	UserEmail:    "geroge@mail.edu",
	//	Amount:       decimal.NewFromFloat(20.11),
	//	CreationDate: "",
	//	UpdateDate:   "",
	//	Status:       "New",
	//}
	//DBServer = t
	http.HandleFunc("/createTransaction", rustServer.createTransaction)
	return http.ListenAndServe(":9000", nil)
}
