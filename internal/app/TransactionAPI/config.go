package TransactionAPI

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func (c *Config) NewConfig() {
	f, err := os.Open("configs/transactionapi.yaml")
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
	fields, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(fields, &c)
	if err != nil {
		log.Fatal(err)
	}
}
