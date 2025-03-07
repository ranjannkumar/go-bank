package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account,error)
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


//This method is used to initialize the database by calling createAccountTable().
func (s *PostgresStore)Init()error{
	return s.createAccountTable()
}

func (s *PostgresStore)createAccountTable()error{
	query :=`create table if not exists account (
	   id serial primary key,
		 first_name varchar(50),
		 last_name varchar(50),
		 number serial,
		 balance serial,
		 created_at timestamp
	)`
	_,err:= s.db.Exec(query)
	return err
}

func (s *PostgresStore)CreateAccount(acc *Account)error{
	  query := `insert into account
		(first_name,last_name,number,balance,created_at)
		values ($1, $2, $3, $4, $5)`

		resp,err := s.db.Query(
			query,
			acc.FirstName,
			acc.LastName,
			acc.Number,
			acc.Balance,acc.CreatedAt,
		)
		if err !=nil {
			return err
		}
		fmt.Printf("%+v\n",resp)
	  
		return nil
}

func (s *PostgresStore)UpdateAccount(*Account)error{
	return nil
}

func (s *PostgresStore)DeleteAccount(id int)error{
	return nil
}

func (s *PostgresStore)GetAccountByID(id int)(*Account,error){
	rows,err := s.db.Query("select * from account where id = $1",id)
	if err!=nil{
		return nil,err
	}
	for rows.Next(){
		return scanIntoAccount(rows)
	}
	return nil,fmt.Errorf("account %d not found",id)
}

func (s *PostgresStore)GetAccounts()([]*Account,error){
	rows,err := s.db.Query("select * from account")
	if err !=nil{
		return nil,err
	}
	accounts := []*Account{}
	for rows.Next(){
		account,err :=scanIntoAccount(rows)
		if err!=nil{
			return nil,err
		}
		accounts = append(accounts, account)
	}
	return accounts,nil
}

func scanIntoAccount(rows *sql.Rows)(*Account,error){
	account := new(Account)
	err:= rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.Balance,
		&account.CreatedAt,
	)
	return account,err

}