package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error
var wg sync.WaitGroup

func init() {
	db, err = sql.Open("postgres", "postgres://postgres:postgres@172.16.238.5/postgres?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	generateData(1, 500000)
}

func hash(secret, data []byte) []byte {
	h := hmac.New(sha256.New, secret)
	h.Write(data)
	return h.Sum(nil)
}
