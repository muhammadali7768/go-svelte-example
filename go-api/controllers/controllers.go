package controllers

import (
	"database/sql"
	"encoding/json"
	"example/books-api/models"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"messge,omitempty"`
}

func createConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected successfully with pg")
	return db
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	insertID := insertBook(book)

	res := response{
		ID:      insertID,
		Message: "Book created successfully",
	}

	json.NewEncoder(w).Encode(res)
}
func GetAllBooks(w http.ResponseWriter, r *http.Request) {}
func GetBookById(w http.ResponseWriter, r *http.Request) {}
func UpdateBook(w http.ResponseWriter, r *http.Request)  {}
func DeleteBook(w http.ResponseWriter, r *http.Request)  {}

func insertBook(book models.Book) int64 {
	db := createConnection()

	defer db.Close()

	sqlStatement := "INSERT INTO books(name,price,publisher) VALUES($1,$2,$3) RETURNING id"
	var id int64
	err := db.QueryRow(sqlStatement, book.Name, book.Price, book.Publisher).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}

	fmt.Printf("inserted a single record with id %v", id)

	return id

}
