package TransactionAPI

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func NewConfig() *Config {
	var config Config
	f, err := os.Open("configs/transactionapi.yaml")
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	flds, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(flds, &config)
	if err != nil {
		log.Fatal(err)
	}
	return &config
}
