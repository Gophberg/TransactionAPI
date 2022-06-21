package TransactionAPI

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func ConnDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password='postgres' dbname=transactions sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, err
}

type DBQuerier interface {
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
