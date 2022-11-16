package store

import (
	_ "github.com/lib/pq"
	"log"
)

func QueryAddMoney(id int, money int) bool {
	db := ConnectDataBase()

	rows, err := db.Query("UPDATE users SET balance = balance + $2 WHERE id = $1;", id, money)

	if err != nil {
		log.Println("[err] Query Adding Money error:", err)
	}

	defer rows.Close()

	return true
}

func QueryReserveMoney(id int, money int) bool {
	db := ConnectDataBase()

	if QueryCheckBalance(id) < money {
		return false
	}

	rows, err := db.Query("UPDATE users SET balance = balance - $2 WHERE id = $1;", id, money)

	_, err = db.Query("UPDATE users SET reserved = $2 WHERE id = $1;", id, money)

	if err != nil {
		log.Println("[err] Query Reserve Money error:", err)
	}

	defer rows.Close()

	return true
}

func QueryTransferMoney(id int, id2 int, money int) bool {
	db := ConnectDataBase()

	if QueryReserveMoney(id, money) {
		rows, err := db.Query("UPDATE users SET reserved = reserved - $2 WHERE id = $1;", id, money)

		if err != nil {
			log.Println("[err] Query Transfer Money error:", err)
		}

		QueryAddMoney(id2, money)
		defer rows.Close()
		return true
	}

	return false
}

func QueryCheckBalance(id int) int {
	db := ConnectDataBase()

	var uidCurrent Users

	uidCurrent.Id = id

	checkBalance, err := db.Query("SELECT balance FROM users WHERE id = $1", uidCurrent.Id)
	if err != nil {
		log.Println("[err] Query Check Balance error: ", err)
	}

	for checkBalance.Next() {
		if err := checkBalance.Scan(&uidCurrent.Balance); err != nil {
			log.Println("[err] Check Balance error:", err)
		}
	}

	return uidCurrent.Balance
}
