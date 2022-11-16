package store

import "time"

type Users struct {
	Id         int    `json:"Id"`
	Name       string `json:"Name"`
	CardNumber string `json:"CardNumber"`
	Balance    int    `json:"Balance"`
	Reserved   int    `json:"Reserved"`
}

type Operations struct {
	OperationId     int       `json:"OperationId"`
	SenderId        int       `json:"SenderId"`
	ReceiverId      int       `json:"ReceiverId"`
	Money           int       `json:"Money"`
	TransactionTime time.Time `json:"TransactionTime"`
}
