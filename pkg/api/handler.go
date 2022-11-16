package api

import (
	"balance-api/pkg/store"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func RouteHandlers() {
	http.HandleFunc("/billing-api/AddMoney", AddMoneyHandler)
	http.HandleFunc("/billing-api/ReserveMoney", ReserveMoneyHandler)
	http.HandleFunc("/billing-api/TransferMoney", TransferMoneyHandler)
	http.HandleFunc("/billing-api/CheckBalance", CheckBalanceHandler)
}

func AddMoneyHandler(writer http.ResponseWriter, request *http.Request) {
	var JsonCome JsonCome

	err := json.NewDecoder(request.Body).Decode(&JsonCome)

	if err != nil {
		writer.Write([]byte("[err] Invalid Json.\n"))
	} else {
		if store.QueryAddMoney(JsonCome.Id, JsonCome.Money) {
			_, errWrite := writer.Write([]byte("Success! Money added!\n"))
			if errWrite != nil {
				fmt.Print("[err] Cannot write response.\n")
			}
		}
	}

}

func ReserveMoneyHandler(writer http.ResponseWriter, request *http.Request) {
	var JsonCome JsonCome

	err := json.NewDecoder(request.Body).Decode(&JsonCome)
	if err != nil {
		writer.Write([]byte("[err] Invalid Json.\n"))
	} else {
		if store.QueryReserveMoney(JsonCome.Id, JsonCome.Money) {
			_, errWrite := writer.Write([]byte("Success! Money is reserved!\n"))
			if errWrite != nil {
				fmt.Print("[err] Cannot write response.\n")
			}
		} else {
			_, errWrite := writer.Write([]byte("Error! Not enough money!\n"))
			if errWrite != nil {
				fmt.Print("[err] Cannot write response.\n")
			}
		}
	}

}

func TransferMoneyHandler(writer http.ResponseWriter, request *http.Request) {
	var JsonTransfer JsonTransfer

	err := json.NewDecoder(request.Body).Decode(&JsonTransfer)
	if err != nil {
		writer.Write([]byte("[err] Invalid Json.\n"))
	} else {
		if store.QueryTransferMoney(JsonTransfer.SenderId, JsonTransfer.ReceiverId, JsonTransfer.Money) {
			_, errWrite := writer.Write([]byte("Success! Money transferred!\n"))
			if errWrite != nil {
				fmt.Print("[err] Cannot write response.\n")
			}
		} else {
			_, errWrite := writer.Write([]byte("Error! Not enough money!\n"))
			if errWrite != nil {
				fmt.Print("[err] Cannot write response.\n")
			}
		}
	}

}

func CheckBalanceHandler(writer http.ResponseWriter, request *http.Request) {
	var JsonCome JsonCome

	err := json.NewDecoder(request.Body).Decode(&JsonCome)
	if err != nil {
		writer.Write([]byte("[err] Invalid Json.\n"))
	} else {
		_, errWrite := writer.Write([]byte(strconv.Itoa(store.QueryCheckBalance(JsonCome.Id))))
		if errWrite != nil {
			fmt.Print("[err] Cannot write response.\n")
		}
	}
}
