package TransactionAPI

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
	Id        int
	UserID    int
	UserEmail string
	//Amount       decimal.Decimal
	Currency     string
	CreationDate int64
	UpdateDate   int64
	Status       Status
}

type Status struct {
	New      string `json:"new"`
	Success  string `json:"success"`
	Failure  string `json:"failure"`
	Error    string `json:"error"`
	Canceled string `json:"canceled"`
}
