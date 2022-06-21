package TransactionAPI

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var DBServer DBQuerier

type RestServer interface {
	getStatusById(http.ResponseWriter, *http.Request)
}

func (t Transaction) getStatusById(w http.ResponseWriter, req *http.Request) {
	splitRoute := strings.Split(req.URL.String(), "?id=")
	id, err := strconv.Atoi(splitRoute[1])
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Value of id: %d\n", id)
	DBServer = t
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

func Start(config *Config) error {
	//DBServer = t
	//allPaymentsByUserId, err := DBServer.GetAllPaymentsByUserId(1)
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println(allPaymentsByUserId)

	// http server
	var rustServer RestServer
	var DBServer = Transaction{}
	rustServer = DBServer
	http.HandleFunc("/getStatus", rustServer.getStatusById)
	return http.ListenAndServe(":9000", nil)
}
