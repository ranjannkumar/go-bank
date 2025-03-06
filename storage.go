package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
}

// This struct represents a PostgreSQL storage layer in your application.
// It holds a database connection (db *sql.DB), which is used to interact with the PostgreSQL database.
type PostgresStore struct {
	db *sql.DB
}

//This function creates and initializes a PostgresStore instance, 
// which is used to interact with a PostgreSQL database.
func NewPostgresStore()(*PostgresStore,error){
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)// Opens a connection pool to the database.
	if err!=nil{
		return nil,err
	}
	if err := db.Ping();err!=nil{
		return nil,err
	}
	//db.Ping() → Sends a test query to check if the database is actually reachable.

	return &PostgresStore{
		db: db,
	},nil
}