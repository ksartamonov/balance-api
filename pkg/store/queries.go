package store

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func QueryAddMoney(id int, money int) bool {
	db := ConnectDataBase()

	rows1, err1 := db.Query("UPDATE users SET balance = balance + $2 WHERE id = $1;", id, money)

	if err1 != nil {
		log.Println("[err] Query Adding Money error.")
	}

	rows2, err2 := db.Query("INSERT INTO operations (ReceiverId, Money, TransactionTime) VALUES ($1, $2, $3)", id, money, time.Now())

	if err2 != nil {
		log.Println("[err] Query Adding Money error.")
	}

	defer rows1.Close()
	defer rows2.Close()

	return true
}

func QueryReserveMoney(id int, money int) bool {
	db := ConnectDataBase()

	if QueryCheckBalance(id) < money {
		return false
	}

	rows1, err1 := db.Query("UPDATE users SET balance = balance - $2 WHERE id = $1;", id, money)

	rows1, err2 := db.Query("UPDATE users SET reserved = reserved + $2 WHERE id = $1;", id, money)

	if err1 != nil || err2 != nil {
		log.Println("[err] Query Reserve Money error.")
	}

	rows2, err3 := db.Query("INSERT INTO operations (SenderId, Money, TransactionTime) VALUES ($1, $2, $3)", id, money, time.Now())

	if err3 != nil {
		log.Println("[err] Query Adding Money error.")
	}

	defer rows1.Close()
	defer rows2.Close()

	return true
}

func QueryTransferMoney(id int, id2 int, money int) bool {
	db := ConnectDataBase()

	if QueryCheckBalance(id) < money {
		return false
	}

	rows1, err1 := db.Query("UPDATE users SET balance = balance - $2 WHERE id = $1;", id, money)
	rows1, err2 := db.Query("UPDATE users SET balance = balance + $2 WHERE id = $1;", id2, money)
	rows2, err3 := db.Query("INSERT INTO operations (SenderId, ReceiverId, Money, TransactionTime) VALUES ($1, $2, $3, $4)", id, id2, money, time.Now())

	if err1 != nil || err2 != nil || err3 != nil {
		log.Println("[err] Query Reserve Money error.")
	}

	defer rows1.Close()
	defer rows2.Close()

	return true
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

// version with writing to file
//func QueryGetUserReport(id int, month int, year int) string {
//	db := ConnectDataBase()
//
//	rows, _ := db.Query(
//		"SELECT CASE WHEN SenderId = 0 THEN 'Adding money' WHEN ReceiverId = 0 THEN 'Reserving money' WHEN ReceiverId = $1 THEN 'Incoming Transfer' WHEN SenderId = $1 THEN 'Outgoing Transfer' ELSE 'Unknown operation' END as \"Operation\", CASE WHEN SenderId = $1 AND ReceiverId != 0 THEN ReceiverId WHEN ReceiverId = $1 AND SenderId != 0 THEN SenderId ELSE 0 END as \"To/From\", money as \"Amount\" FROM operations WHERE (ReceiverId = $1 OR SenderId = $1) AND ($2 = date_part('month', transactiontime) AND $3 = date_part('year', transactiontime));", id, month, year)
//	reportName := "(user)(" + strconv.Itoa(month) + "." + strconv.Itoa(year) + ")" + "report(id=" + strconv.Itoa(id) + ").csv"
//	err := sqltocsv.WriteFile(reportName, rows)
//	if err != nil {
//		log.Println("[err] Query Get User Report error: ", err)
//	}
//	defer rows.Close()
//
//	return reportName
//}

func QueryGetUserReport(id int, month int, year int) *sql.Rows {
	db := ConnectDataBase()

	rows, err := db.Query(
		"SELECT OperationId, CASE WHEN SenderId = 0 THEN 'Adding money' WHEN ReceiverId = 0 THEN 'Reserving money' WHEN ReceiverId = $1 THEN 'Incoming Transfer' WHEN SenderId = $1 THEN 'Outgoing Transfer' ELSE 'Unknown operation' END as \"Operation\", CASE WHEN SenderId = $1 AND ReceiverId != 0 THEN ReceiverId WHEN ReceiverId = $1 AND SenderId != 0 THEN SenderId ELSE 0 END as \"To/From\", money as \"Amount\" FROM operations WHERE (ReceiverId = $1 OR SenderId = $1) AND ($2 = date_part('month', transactiontime) AND $3 = date_part('year', transactiontime));", id, month, year)
	defer rows.Close()
	if err != nil {
		log.Println("[err] Query Get User Report error: ", err)
	}
	return rows
}
