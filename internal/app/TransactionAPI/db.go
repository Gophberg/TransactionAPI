package TransactionAPI

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConnDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password='postgres' dbname=transactions sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, err
}
