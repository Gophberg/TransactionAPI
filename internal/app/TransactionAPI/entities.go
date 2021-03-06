package TransactionAPI

import (
	"github.com/shopspring/decimal"
)

type Config struct {
	Jwttoken     string `yaml:"jwttoken"`
	Dbhost       string `yaml:"dbhost"`
	Dbname       string `yaml:"dbname"`
	Dbusername   string `yaml:"dbusername"`
	Dbpassword   string `yaml:"dbpassword"`
	Dockerdbport string `yaml:"dockerdbport"`
}

type Transaction struct {
	Id           int64           `json:"id"`
	UserID       int64           `json:"userid"`
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
