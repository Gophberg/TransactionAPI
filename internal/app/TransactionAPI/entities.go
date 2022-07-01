package TransactionAPI

import (
	"github.com/shopspring/decimal"
)

//const (
//	DB_HOST     = config.Dbhost
//	DB_USERNAME = config.Dbusername
//	DB_PASSWORD = config.Dbpassword
//	DB_NAME     = config.Dbname
//)

type Config struct {
	Jwttoken     string `yaml:"jwttoken"`
	Dbhost       string `yaml:"dbhost"`
	Dbname       string `yaml:"dbname"`
	Dbusername   string `yaml:"dbusername"`
	Dbpassword   string `yaml:"dbpassword"`
	Dockerdbport string `yaml:"dockerdbport"`
}

type User struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

type Transaction struct {
	Id           int             `json:"id"`
	UserID       int             `json:"userid"`
	UserEmail    string          `json:"useremail"`
	Amount       decimal.Decimal `json:"amount"`
	Currency     string          `json:"currency"`
	CreationDate string          `json:"datecreated"`
	UpdateDate   string          `json:"dateupdated"`
	Status       string          `json:"status"`
}

type Status struct {
	New      string `json:"new"`
	Success  string `json:"success"`
	Failure  string `json:"failure"`
	Error    string `json:"error"`
	Canceled string `json:"canceled"`
}

//type DBQuerier interface {
//	CreateTransaction(Config) error
//	GetTransactionStatusById(Config, int) (string, error)
//	GetAllTransactionsByUserId(Config, int) ([]Transaction, error)
//	GetAllTransactionsByUserEmail(Config, string) ([]Transaction, error)
//}
//
//type RestServer interface {
//	createTransaction(http.ResponseWriter, *http.Request)
//	getStatusById(http.ResponseWriter, *http.Request)
//	getAllTransactionsByUserId(http.ResponseWriter, *http.Request)
//	getAllTransactionsByUserEmail(http.ResponseWriter, *http.Request)
//}

type restTransaction struct {
	transaction Transaction
}

type dbTransaction struct {
	transaction Transaction
}
