package controllers

import (
	"OnlineBank/database"
	"OnlineBank/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jackc/pgx"
	"log"
	"net/http"
	"strconv"
)

func TxAccHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var request models.AccTxReq
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("Ошибка при получении данных:", err.Error())
		w.WriteHeader(400)
		ErrorMessage(w, err)
		return
	}
	log.Println("request:", request)

	id, err := strconv.Atoi(r.Header.Get("user_id"))
	if err != nil {
		log.Println("User_id не конвертируется:", err.Error())
		return
	}
	log.Println("user_id:", id)

	receiverAcc, receiverBal := getReceiverData(w, request.ReceiverAcc)
	log.Println("receiverAcc:", receiverAcc, "receiverBal:", receiverBal)
	senderAcc, senderBal := getSenderData(w, id)
	log.Println("senderAcc:", senderAcc, "senderBal:", senderBal)

	tx, err := database.GetDBConn.Begin()
	if err != nil {
		log.Println("не удалось начать транзакцию по причине:", err.Error())
		w.WriteHeader(500)
		ErrorMessage(w, err)
		return
	}

	if request.Amount > senderBal {
		err = errors.New("не достаточно денег на счету")
		log.Println("Ошибка баланса:", err.Error())
		ErrorMessage(w, err)
		return
	}

	defer func() {
		log.Println("check defer")
		if err != nil {
			err = tx.Rollback()
			log.Println("Ошибка при откате транзакции по причине:", err.Error())
			return
		}
		err = tx.Commit()
		if err != nil {
			log.Println("не удалось создать транзакцию по причине:", err.Error())
			return
		}
	}()

	updateSendBal(w, tx, senderBal, request.Amount, id)
	updateRecBal(w, tx, receiverBal, request.Amount, receiverAcc)
	txExec(w, tx, senderAcc, receiverAcc, request.Amount)

	responseBody := models.ResponseStruct{
		Message:     fmt.Sprintf("Перевод в размере %v у.е был успешно перечислен!", request.Amount),
		Error:       "0",
		Description: "Операция успешно завершена",
	}
	err = json.NewEncoder(w).Encode(&responseBody)
	if err != nil {
		log.Println("Ошибка в функции перевода денег:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		ErrorMessage(w, err)
		return
	}

}

//getReceiverData - returns receiver data/ возвращает данные получателя
func getReceiverData(w http.ResponseWriter, account string) (string, int) {
	var receiver models.ClientAccounts
	err := database.GetDBConn.QueryRow(`select acc_number, balance from client_accounts where acc_number = $1`, account).Scan(&receiver.AccNumber, &receiver.Balance)
	if err != nil {
		log.Println("не удалось получить данные получателя с БД, ошибка:", err.Error())
		ErrorMessage(w, err)
		return "", 0
	}
	return receiver.AccNumber, receiver.Balance
}

//getSenderData - returns sender data/возвращает данные получателя
func getSenderData(w http.ResponseWriter, id int) (string, int) {
	var sender models.ClientAccounts
	err := database.GetDBConn.QueryRow(`select acc_number, balance from client_accounts where client_id = $1`, id).Scan(&sender.AccNumber, &sender.Balance)
	if err != nil {
		log.Println("не удалось получить данные отправителя с БД, ошибка:", err.Error())
		ErrorMessage(w, err)
		return "", 0
	}
	return sender.AccNumber, sender.Balance
}

//updateRecBal - updates receiver balance during transaction/обновляет баланс получателя при транзакции
func updateRecBal(w http.ResponseWriter, tx *pgx.Tx, balance, amount int, account string) {
	_, err := tx.Exec(`update client_accounts set balance = $1 where acc_number = $2`, balance+amount, account)
	if err != nil {
		log.Println("Ошибка при изменении счета получателя по причине:", err.Error())
		w.WriteHeader(500)
		ErrorMessage(w, err)
		return
	}
}

//updateSendBal - updates sender balance during transactions/бновляет баланс отправителя при транзакции
func updateSendBal(w http.ResponseWriter, tx *pgx.Tx, balance, amount int, id int) {
	_, err := tx.Exec(`update client_accounts set balance = $1 where client_id = $2`, balance-amount, id)
	if err != nil {
		log.Println("Ошибка при изменении счета отправителя по причине:", err.Error())
		w.WriteHeader(500)
		ErrorMessage(w, err)
		return
	}
}

//txExec - writes transaction data to DB/записывает данные о транзакции в БД
func txExec(w http.ResponseWriter, tx *pgx.Tx, senderAcc, recAcc string, amount int) {
	_, err := tx.Exec(`insert into transactions (payer_acc, beneficiary_acc, transfer_amount) values ($1, $2, $3)`, senderAcc, recAcc, amount)
	if err != nil {
		fmt.Println("Ошибка при записи транзакции в БД по причине:", err)
		w.WriteHeader(500)
		ErrorMessage(w, err)
		return
	}
}
