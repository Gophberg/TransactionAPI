package TransactionAPI

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func ConnDB() (*sql.DB, error) {
	url := fmt.Sprintf("host=%v user=%v password='%v' dbname=%v sslmode=disable", config.Dbhost, config.Dbusername, config.Dbpassword, config.Dbname)
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return db, err
}

func (t *Transaction) CreateTransaction(c Transaction) (int, error) {
	db, err := ConnDB()
	if err != nil {
		return 0, err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	log.Printf("Reseived Credentials: %v", c)

	err = db.QueryRow(
		`INSERT INTO transactions (userid, useremail, amount, currency, creationdate, updatedate, status) 
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		c.UserID,
		c.UserEmail,
		c.Amount,
		c.Currency,
		c.CreationDate,
		c.UpdateDate,
		c.Status,
	).Scan(&t.Id)
	return t.Id, err
}

func (t Transaction) GetTransactionStatusById(c Transaction) (string, error) {
	db, err := ConnDB()
	if err != nil {
		return "", err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	log.Printf("Query request <%v> status\n", c.Id)

	query := fmt.Sprintf(`SELECT status FROM transactions WHERE id = %d;`, c.Id)
	if err := db.QueryRow(query).Scan(&t.Status); err != nil {
		if err == sql.ErrNoRows {
			return "", err
		}
	}
	fmt.Printf("Status is <%v>\n", t.Status)
	return t.Status, nil
}

func (t Transaction) GetAllTransactionsByUserId(c Transaction) ([]Transaction, error) {
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

	query := fmt.Sprintf(`SELECT * FROM transactions WHERE userid = %d;`, c.UserID)

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

func (t Transaction) GetAllTransactionsByUserEmail(c Transaction) ([]Transaction, error) {
	db, err := ConnDB()
	if err != nil {
		return nil, err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	fmt.Printf("Requested query email: %v\n", c.UserEmail)
	query := fmt.Sprintf(`SELECT * FROM transactions WHERE useremail = '%v';`, c.UserEmail)

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
