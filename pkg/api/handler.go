package api

import (
	"balance-api/pkg/api/dto"
	"balance-api/pkg/store"
	"encoding/json"
	"fmt"
	"github.com/joho/sqltocsv"
	"log"
	"net/http"
)

func RouteHandlers() {
	http.HandleFunc("/billing-api/AddMoney", AddMoneyHandler)
	http.HandleFunc("/billing-api/ReserveMoney", ReserveMoneyHandler)
	http.HandleFunc("/billing-api/TransferMoney", TransferMoneyHandler)
	http.HandleFunc("/billing-api/CheckBalance", CheckBalanceHandler)
	http.HandleFunc("/billing-api/GetUserReport", GetUserReportHandler)
}

func AddMoneyHandler(writer http.ResponseWriter, request *http.Request) {
	var JsonCome dto.JsonCome

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
	var JsonCome dto.JsonCome

	err := json.NewDecoder(request.Body).Decode(&JsonCome)
	if err != nil {
		writer.Write([]byte("[err] Invalid Json.\n"))
	} else {
		if store.QueryReserveMoney(JsonCome.Id, JsonCome.Money) {
			_, errWrite := writer.Write([]byte("Success! Money reserved!\n"))
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
	var JsonTransfer dto.JsonTransfer

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

	var JsonCome dto.JsonCome

	err := json.NewDecoder(request.Body).Decode(&JsonCome)
	if err != nil {
		writer.Write([]byte("[err] Invalid Json.\n"))
	} else {
		writer.Header().Set("Content-Type", "application/json")
		resp := make(map[string]int)
		resp["id"] = JsonCome.Id
		resp["balance"] = store.QueryCheckBalance(JsonCome.Id)
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("[err] JSON marshal error:%s", err)
		}
		writer.Write(jsonResp)
	}
}

// Version with writing to file
//func GetUserReportHandler(writer http.ResponseWriter, request *http.Request) {
//	var JsonReport JsonReport
//
//	err := json.NewDecoder(request.Body).Decode(&JsonReport)
//	if err != nil {
//		writer.Write([]byte("[err] Invalid Json.\n"))
//	} else {
//		writer.Write([]byte("Report saved in file " + store.QueryGetUserReport(JsonReport.Id, JsonReport.Month, JsonReport.Year)))
//	}
//}

func GetUserReportHandler(writer http.ResponseWriter, request *http.Request) {
	var JsonReport dto.JsonReport
	err := json.NewDecoder(request.Body).Decode(&JsonReport)
	if err != nil {
		writer.Write([]byte("[err] Invalid Json.\n"))
	} else {
		http.HandleFunc("/", uploadUserReportHelper)
		http.ListenAndServe(":9191", nil)
		sqltocsv.Write(writer, store.QueryGetUserReport(JsonReport.Id, JsonReport.Month, JsonReport.Year))
	}
}

func uploadUserReportHelper(writer http.ResponseWriter, request *http.Request) {
	//writer.Write([]byte("Report saved in file " + store.QueryGetUserReport(JsonReport.Id, JsonReport.Month, JsonReport.Year)))
	writer.Header().Set("Content-type", "text/csv")
	writer.Header().Set("Content-Disposition", "attachment; filename=\"report.csv\"")
}
