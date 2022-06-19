package TransactionAPI

import (
	"database/sql"
	"fmt"
	"log"
)

type DbServerInterface interface {
	CheckStatusById(int) (string, error)
	GetAllPaymentsByUserId(int) ([]Transaction, error)
}

func (t Transaction) CheckStatusById(id int) (string, error) {
	db, err := ConnDB()
	if err != nil {
		return "", err
	}
	defer db.Close()
	query := fmt.Sprintf(`SELECT status FROM transactions WHERE id = %d;`, id)
	if err := db.QueryRow(query).Scan(&t.Status); err != nil {
		if err == sql.ErrNoRows {
			return "", err
		}
	}
	return t.Status, nil
}

func (t Transaction) GetAllPaymentsByUserId(id int) ([]Transaction, error) {
	var items []Transaction
	db, err := ConnDB()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf(`SELECT * FROM transactions WHERE userid = %d;`, id)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.Id,
			&i.UserID,
			&i.UserEmail,
			&i.Amount,
			&i.Currency,
			&i.CreationDate,
			&i.UpdateDate,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, nil
}

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
	return nil
}
