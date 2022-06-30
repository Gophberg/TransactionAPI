package TransactionAPI

import (
	"net/http"
)

func Start(config *Config) error {
	var restServer RestServer
	var DBServer = Transaction{}
	restServer = DBServer
	http.HandleFunc("/getStatus", restServer.getStatusById)
	http.HandleFunc("/getAllTransactionsByUserId", restServer.getAllTransactionsByUserId)
	http.HandleFunc("/getAllTransactionsByUserEmail", restServer.getAllTransactionsByUserEmail)
	http.HandleFunc("/createTransaction", restServer.createTransaction)
	return http.ListenAndServe(":9000", nil)
}
