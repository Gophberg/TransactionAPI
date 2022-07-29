package TransactionAPI

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func ConnDB() (*sql.DB, error) {
	url := fmt.Sprintf("host=%v user=%v password='%v' dbname=%v sslmode=disable", config.Dbhost, config.Dbusername, config.Dbpassword, config.Dbname)
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return db, err
}

func (t *Transaction) createRecord(c Transaction) (int64, error) {
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

	log.Printf("[DB] Reseived createRecord Credentials: %v", c)

	c.CreationDate = time.Now().Format(time.RFC3339)
	c.UpdateDate = ""

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

func (t *Transaction) updateRecord(c Transaction) (int64, error) {
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

	log.Printf("[DB] Reseived updateRecord Credentials: %v", c)

	c.UpdateDate = time.Now().Format(time.RFC3339)

	sqlStatements := `UPDATE transactions SET updatedate = $1, status = $2 WHERE id = $3;`
	_, err = db.Exec(sqlStatements, c.UpdateDate, c.Status, c.Id)
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func (t Transaction) readRecord(c Transaction) (string, error) {
	db, err := ConnDB()
	if err != nil {
		return "", err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	log.Printf("[DB] Query request <%v> status\n", c.Id)

	query := fmt.Sprintf(`SELECT status FROM transactions WHERE id = %d;`, c.Id)
	if err := db.QueryRow(query).Scan(&t.Status); err != nil {
		if err == sql.ErrNoRows {
			return "", err
		}
	}
	log.Printf("[DB] Status is <%v>\n", t.Status)
	return t.Status, nil
}

func (t Transaction) readRecords(c Transaction, urlPath string) ([]Transaction, error) {
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

	switch {
	//case urlPath == "/getTransactionStatusById":

	case urlPath == "/getAllTransactionsByUserId":
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
	case urlPath == "/getAllTransactionsByUserEmail":
		query := fmt.Sprintf(`SELECT * FROM transactions WHERE useremail = '%v';`, c.UserEmail)
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
	return nil, nil
}
