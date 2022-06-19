package TransactionAPI

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type DbServerInterface interface {
	CheckStatusById(Config, int) (string, error)
	GetAllPaymentsByUserId(Config, int) ([]Transaction, error)
	CreatePayment(transaction Transaction) error
}

func ConnDB(config Config) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("host=%v user=%v password='%v' dbname=transactions sslmode=disable", config.DatabaseHost, config.Dbusername, config.Dbpassword)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	return db, err
}

type TransactionsRepository struct {
	db *sql.DB
}

func (t Transaction) CheckStatusById(config Config, id int) (string, error) {
	db, err := ConnDB(config)
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

func (t Transaction) GetAllPaymentsByUserId(config Config, id int) ([]Transaction, error) {
	var items []Transaction
	db, err := ConnDB(config)
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

///////////////////////////////////////////////////////////////////////////////////////////

//func (t TransactionsRepository) CreatePayment(transaction Transaction) error {
//	t.db, err := ConnDB()
//	if err != nil {
//		return err
//	}
//	tr := Transaction{
//		UserID: 4,
//
//	}
//	defer t.db.Close()
//	return t.db.QueryRow(`INSERT INTO transactions (userid, useremail, amount, currency, creationdate, updatedate, status) VALUES (tr.UserId, tr.User, 321.11, 'USD', TIMESTAMP '2022-06-19 21:00:06', TIMESTAMP '2022-06-19 21:00:07', 'Success'`).Scan(t.Id)
//}
