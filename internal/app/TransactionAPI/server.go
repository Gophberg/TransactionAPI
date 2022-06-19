package TransactionAPI

import (
	"fmt"
	"log"
)

func Start(config *Config) error {
	var DBServer DbServerInterface
	transactionImplementation := Transaction{
		UserID: 1,
	}
	DBServer = transactionImplementation
	transactionStatus, err := DBServer.CheckStatusById(1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(transactionStatus)
	allPaymentsByUserId, err := DBServer.GetAllPaymentsByUserId(1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(allPaymentsByUserId)

	//DBServer.CreatePayment()
	
	return nil
}
