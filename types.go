package main

import (
	"math/rand"
	"time"
)

type TransferRequest struct{
	ToAccount   int   `json:"toAccount"`
	Amount      int   `json:"amount"`
}

type CreateAccountRequest struct{
	FirstName  string  `json:"firstName"`
	LastName   string  `json:"lastName"`
}

type Account struct {
	ID        int      `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewAccount(FirstName, LastName string) *Account {
	return &Account{
		// ID:				 rand.Intn(10000),
		FirstName: FirstName,
		LastName:	 LastName,
		Number: 		int64(rand.Intn(1000000)),
		CreatedAt: time.Now().UTC(),
	}
}