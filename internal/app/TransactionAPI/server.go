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
var cancelTrChan = make(chan int64, 10)

func Start() error {
	config.NewConfig()
	t := Transaction{}
	http.HandleFunc("/getTransactionStatusById", t.getTransactionStatusById)
	http.HandleFunc("/getAllTransactionsByUserId", t.getAllTransactions)
	http.HandleFunc("/getAllTransactionsByUserEmail", t.getAllTransactions)
	http.HandleFunc("/createTransaction", t.createTransaction)
	http.HandleFunc("/cancelTransaction", t.cancelTransaction)
	return http.ListenAndServe(":9000", nil)
}
