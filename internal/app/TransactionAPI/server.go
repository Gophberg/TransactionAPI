package TransactionAPI

import (
	"net/http"
)

var config Config

func Start() error {
	config.NewConfig()
	t := Transaction{}
	//var restServer RestServer
	//var DBServer = Transaction{}
	//restServer = DBServer
	http.HandleFunc("/getTransactionStatusById", t.getTransactionStatusById)
	http.HandleFunc("/getAllTransactionsByUserId", t.getAllTransactionsByUserId)
	http.HandleFunc("/getAllTransactionsByUserEmail", t.getAllTransactionsByUserEmail)
	http.HandleFunc("/createTransaction", t.createTransaction)
	return http.ListenAndServe(":9000", nil)
}
