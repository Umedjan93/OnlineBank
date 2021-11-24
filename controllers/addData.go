package controllers

import (
	"OnlineBank/database"
	"OnlineBank/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//AddClientHandler - добавляет нового пользователя (физ.лица) в БД/ adds new user for phys.body
func AddClientHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var newClient models.ClientsData

	err := json.NewDecoder(r.Body).Decode(&newClient)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ErrorMessage(w, err)
		return
	}

	_, err = database.GetDBConn.Exec(database.AddClientQuery, newClient.Name, newClient.Phone, newClient.Login, newClient.Password)
	if err != nil {
		log.Println("ошибка при добавлении клиента в БД:", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		ErrorMessage(w, err)
		return
	}
	responseBody := models.ResponseStruct{
		Message:     fmt.Sprintf("Клиент %v был успешно добавлен", newClient.Name),
		Error:       "0",
		Description: "Операция успешно завершена",
	}
	err = json.NewEncoder(w).Encode(&responseBody)
	if err != nil {
		log.Println("Ошибка в функции добавления нового клиента:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		ErrorMessage(w, err)
		return
	}
}

// AddMerchantHandler - добавляет нового пользователя (юр.лица) в БД// adds new merchant user
func AddMerchantHandler(w http.ResponseWriter, r *http.Request) {

	var newMerchant models.MerchantsData
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	err := json.NewDecoder(r.Body).Decode(&newMerchant)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ErrorMessage(w, err)
		return
	}

	_, err = database.GetDBConn.Exec(database.AddMerchantQuery, newMerchant.Name, newMerchant.Phone, newMerchant.Login, newMerchant.Password)
	if err != nil {
		log.Println("ошибка при добавлении компании в БД:", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		ErrorMessage(w, err)
		return
	}
	responseBody := models.ResponseStruct{
		Message:     fmt.Sprintf("Компания %v была успешно добавлена", newMerchant.Name),
		Error:       "0",
		Description: "Операция успешно завершена",
	}
	err = json.NewEncoder(w).Encode(&responseBody)
	if err != nil {
		log.Println("Ошибка в функции добавления новой компании:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		ErrorMessage(w, err)
		return
	}
}

//AddMerchantAccHandler - добавляет новый банковский счет компании/adds merchant account(for companies)
func AddMerchantAccHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var newMerchantAcc models.MerchantAccounts

	err := json.NewDecoder(r.Body).Decode(&newMerchantAcc)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ErrorMessage(w, err)
		return
	}

	_, err = database.GetDBConn.Exec(database.AddMerchantAccQuery, newMerchantAcc.AccNumber, newMerchantAcc.MerchantID, newMerchantAcc.Balance)
	if err != nil {
		log.Println("ошибка при добавлении коммерчиского счета в БД:", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		ErrorMessage(w, err)
		return
	}
	responseBody := models.ResponseStruct{
		Message:     fmt.Sprintf("Счет %v был успешно добавлен", newMerchantAcc.AccNumber),
		Error:       "0",
		Description: "Операция успешно завершена",
	}
	err = json.NewEncoder(w).Encode(&responseBody)
	if err != nil {
		log.Println("Ошибка в функции добавления счета компании:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		ErrorMessage(w, err)
		return
	}
}

//AddClientAccHandler - добавляет новый счет клиенту банка/adds new user account
func AddClientAccHandler(w http.ResponseWriter, r *http.Request) {

	var newClientAcc models.ClientAccounts
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	err := json.NewDecoder(r.Body).Decode(&newClientAcc)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ErrorMessage(w, err)
		return
	}

	_, err = database.GetDBConn.Exec(database.AddClientAccQuery, newClientAcc.AccNumber, newClientAcc.ClientID, newClientAcc.Balance)
	if err != nil {
		log.Println("ошибка при добавлении коммерчиского счета в БД:", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		ErrorMessage(w, err)
		return
	}
	responseBody := models.ResponseStruct{
		Message:     fmt.Sprintf("Счет %v был успешно добавлен", newClientAcc.AccNumber),
		Error:       "0",
		Description: "Операция успешно завершена",
	}
	err = json.NewEncoder(w).Encode(&responseBody)
	if err != nil {
		log.Println("Ошибка в функции добавления счета клиента:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		ErrorMessage(w, err)
		return
	}
}

//AddService - добавляет новый вид онлайн-услуг в БД/adds new online service
func AddService(w http.ResponseWriter, r *http.Request) {

	var newService models.Services
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	err := json.NewDecoder(r.Body).Decode(&newService)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ErrorMessage(w, err)
		return
	}

	_, err = database.GetDBConn.Exec(database.AddServiceQuery, newService.ServiceName, newService.Price, newService.MerchantName, newService.MerchantID)
	if err != nil {
		log.Println("ошибка при добавлении сервиса в БД:", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		ErrorMessage(w, err)
		return
	}
	responseBody := models.ResponseStruct{
		Message:     fmt.Sprintf("Сервис %v был успешно добавлен", newService.ServiceName),
		Error:       "0",
		Description: "Операция успешно завершена",
	}
	err = json.NewEncoder(w).Encode(&responseBody)
	if err != nil {
		log.Println("Ошибка в функции добавления нового севрсиа:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		ErrorMessage(w, err)
		return
	}
}

//AddAtmHandler - добавляет банкомат в БД/ adds new ATM to DB
func AddAtmHandler(w http.ResponseWriter, r *http.Request) {

	var newAtm models.ATMs
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	err := json.NewDecoder(r.Body).Decode(&newAtm)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ErrorMessage(w, err)
		return
	}

	_, err = database.GetDBConn.Exec(database.AddATMQuery, newAtm.Address, newAtm.Balance, newAtm.MaxCashLimit, newAtm.Commission, newAtm.Owner)
	if err != nil {
		log.Println("ошибка при добавлении банкомата в БД:", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		ErrorMessage(w, err)
		return
	}
	responseBody := models.ResponseStruct{
		Message:     fmt.Sprintf("Банкомат на %v был успешно добавлен", newAtm.Address),
		Error:       "0",
		Description: "Операция успешно завершена",
	}
	err = json.NewEncoder(w).Encode(&responseBody)
	if err != nil {
		log.Println("Ошибка в функции добавления нового банкомата:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		ErrorMessage(w, err)
		return
	}
}
