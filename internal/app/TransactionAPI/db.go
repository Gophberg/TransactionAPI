package TransactionAPI

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

//type DBServer struct {
//	connDb *sql.DB
//	transaction Transaction
//}

func ConnDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password='postgres' dbname=transactions sslmode=disable")
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
	//err = db.QueryRow(
	//	`INSERT INTO transactions (userid, useremail, amount, currency, creationdate, updatedate, status)
	//	VALUES (5, 'john@mail.edu', 33.12, 'rub', '2022-06-23T15:55:00Z', '2022-06-23T15:55:01Z', 'new') RETURNING id`).Scan(t.Id)
	return err
	//return db.QueryRow(
	//	`insert into transactions (userid, useremail, amount, currency, creationdate, updatedate, status)
	//	values (4, 'john@mail.edu', 33.12, 'rub', '2022-06-23T15:55:00Z', '2022-06-23T15:55:01Z', 'new') returning id`,)
	//}

	//return r.store.db.QueryRow(
	//	"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
	//	u.Email,
	//	u.EncryptedPassword,
	//).Scan(&u.ID)

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

	query := fmt.Sprintf(`SELECT status FROM transactions WHERE id = %d;`, id)
	if err := db.QueryRow(query).Scan(&t.Status); err != nil {
		if err == sql.ErrNoRows {
			return "", err
		}
	}
	return t.Status, nil
}

func (t Transaction) GetAllTransactionsByUserId(id int) ([]Transaction, error) {
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
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

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
	var items []Transaction
	db, err := ConnDB()
	if err != nil {
		return nil, err
	}

	fmt.Printf("Requested query email: %v\n", email)
	query := fmt.Sprintf(`SELECT * FROM transactions WHERE useremail = '%v';`, email)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

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
