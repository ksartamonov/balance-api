package store

import (
	"balance-api/pkg/config"
	"database/sql"
	"fmt"
	"log"
)

func ConnectDataBase() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		config.GetConfig().DbHost, config.GetConfig().DbPort, config.GetConfig().User,
		config.GetConfig().Password, config.GetConfig().DbName, config.GetConfig().SSLMode)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("[err] Cannot open database.", err)
	}

	err = db.Ping()
	if err != nil {
		log.Print(err)
	}

	return db
}
