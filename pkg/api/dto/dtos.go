package dto

type JsonCome struct {
	Id    int `json:"Id"`
	Money int `json:"Money"`
	//TransactionTime time.Time `json:"TransactionTime"`
}

type JsonTransfer struct {
	SenderId   int `json:"SenderId"`
	ReceiverId int `json:"ReceiverId"`
	Money      int `json:"Money"`
	//TransactionTime time.Time `json:"TransactionTime"`
}

type JsonBalance struct {
	Money int `json:"Money"`
}

type JsonReport struct {
	Id    int `json:"Id"`
	Month int `json:"Month"`
	Year  int `json:"Year"`
}
