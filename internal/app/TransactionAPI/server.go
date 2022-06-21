package TransactionAPI

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Start(config *Config) error {
	var DBServer DbServerInterface
	transactionImplementation := Transaction{
		UserID: 1,
	}
	DBServer = transactionImplementation
	transactionstatus, err := DBServer.CheckStatusById(1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(transactionstatus)
	allPaymentsByUserId, err := DBServer.GetAllPaymentsByUserId(1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(allPaymentsByUserId)

	// http server

	http.HandleFunc("/getStatus", getStatusById)

	return http.ListenAndServe(":9000", nil)

	//return nil
}

func getStatusById(w http.ResponseWriter, req *http.Request) {
	splitRoute := strings.Split(req.URL.String(), "?id=")
	id, err := strconv.Atoi(splitRoute[1])
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Value of id: %d\n", id)
	var DBServer DbServerInterface
	transactionImplenemtation := Transaction{
		UserID: id,
	}
	DBServer = transactionImplenemtation
	status, err := DBServer.CheckStatusById(id)
	if err != nil {
		log.Println(err)
	}
	b, err := fmt.Fprint(w, status)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(b)
}
