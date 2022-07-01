package TransactionAPI

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var config Config

func ConnDB() (*sql.DB, error) {
	url := fmt.Sprintf("host=%v user=$v password='%v' dbname=%v sslmode=disable", config.Dbhost, config.Dbusername, config.Dbpassword, config.Dbname)
	fmt.Printf("Connection db URL <%v> \n", url)
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return db, err
}

func (t Transaction) CreateTransaction() error {
	db, err := ConnDB()
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	fmt.Println(t)
	var tt *int
	err = db.QueryRow(
		`INSERT INTO transactions (userid, useremail, amount, currency, creationdate, updatedate, status) 
		VALUES (5, 'john@mail.edu', 33.12, 'rub', '2022-06-23T15:55:00Z', '2022-06-23T15:55:01Z', 'new') RETURNING id`).Scan(&tt)
	fmt.Println(*tt)
	return err
}

func (t Transaction) GetTransactionStatusById(id int) (string, error) {
	db, err := ConnDB()
	if err != nil {
		return "", err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	fmt.Printf("Query status <%v> id\n", id)
	query := fmt.Sprintf(`SELECT status FROM transactions WHERE id = %d;`, id)
	if err := db.QueryRow(query).Scan(&t.Status); err != nil {
		if err == sql.ErrNoRows {
			return "", err
		}
	}
	fmt.Printf("Status is <%v>\n", t.Status)
	return t.Status, nil
}

func (t Transaction) GetAllTransactionsByUserId(id int) ([]Transaction, error) {
	db, err := ConnDB()
	if err != nil {
		return nil, err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var items []Transaction

	query := fmt.Sprintf(`SELECT * FROM transactions WHERE userid = %d;`, id)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

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

func (t Transaction) GetAllTransactionsByUserEmail(email string) ([]Transaction, error) {
	db, err := ConnDB()
	if err != nil {
		return nil, err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	fmt.Printf("Requested query email: %v\n", email)
	query := fmt.Sprintf(`SELECT * FROM transactions WHERE useremail = '%v';`, email)

	var items []Transaction

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

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
