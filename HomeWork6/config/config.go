package config

import (
	"encoding/json"
	"log"
	"os"
)

type Database struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Ip       string `json:"ip,omitempty"`
	Port     string `json:"port,omitempty"`
	Dbname   string `json:"dbname,omitempty"`
}

func Loading(base *Database) error {
	filePath := "./configs.json"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&base)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil

}
