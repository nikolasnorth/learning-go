package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Account struct {
	Name      string  `json:"Name"`
	AccountId int     `json:"AccountID"`
	Balance   float64 `json:"Balance"`
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("invalid input or missing filepath. Example usage: go run accounts.go /path/to/file")
	}

	filepath := os.Args[1]
	infile, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer infile.Close()

	bytes, err := io.ReadAll(infile)
	if err != nil {
		log.Fatal(err.Error())
	}

	var accounts []Account
	err = json.Unmarshal(bytes, &accounts)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := range accounts {
		accounts[i].Balance += 100
	}

	bytes, err = json.Marshal(accounts)
	if err != nil {
		log.Fatal(err.Error())
	}

	outfile, err := os.Create("accountsUpdated.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer outfile.Close()

	outfile.Write(bytes)
}
