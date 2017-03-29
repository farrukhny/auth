package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://golang:docker17@ec2-54-208-7-163.compute-1.amazonaws.com/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("You connected to your database.")
}

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

func IndexDB() []Book {

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()

	bks := make([]Book, 0)

	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			log.Print(err)
		}
		bks = append(bks, bk)
	}

	return bks

}
