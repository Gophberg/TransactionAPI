package TransactionAPI

import (
	"github.com/shopspring/decimal"
)

type Config struct {
	Jwttoken     string `yaml:"jwttoken"`
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
