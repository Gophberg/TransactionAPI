package TransactionAPI

import (
	"flag"
	"net/http"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

var config Config

func Start() error {
	config.NewConfig()
	t := Transaction{}
	http.HandleFunc("/getTransactionStatusById", t.getTransactionStatusById)
	http.HandleFunc("/getAllTransactionsByUserId", t.getAllTransactionsByUserId)
	http.HandleFunc("/getAllTransactionsByUserEmail", t.getAllTransactionsByUserEmail)
	http.HandleFunc("/createTransaction", t.createTransaction)
	return http.ListenAndServe(":9000", nil)
}
