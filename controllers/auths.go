package controllers
//This file contains all functions and handlers for authorisation procedure
//данный файл содержит все функции и хендлеры для авторизации пользователей

import (
	"OnlineBank/database"
	"OnlineBank/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//AuthClient - client authorisation main func/ основная функция авторизации клиентов
func AuthClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var data models.AuthData
	json.NewDecoder(r.Body).Decode(&data)
	log.Println("data", data)
	passed, userid := CheckClientData(data.Login, data.Password)
	if passed != true {
		w.WriteHeader(400)
		err := json.NewEncoder(w).Encode(models.ErrorDescr{Reason: "Неверный логин или пароль"})
		if err != nil {
			log.Println("Неверные данные:", err.Error())
		}
		return
	}

	token := GenerateToken(userid, "client")
	json.NewEncoder(w).Encode(models.ActiveToken{token})
}

//AuthManager - manager authorisation main func/ основная функция авторизации менеджеров
func AuthManager(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var data models.AuthData
	json.NewDecoder(r.Body).Decode(&data)
	passed, userid := CheckManagerData(data.Login, data.Password)
	if passed != true {
		w.WriteHeader(400)
		err := json.NewEncoder(w).Encode(models.ErrorDescr{Reason: "Неверный логин или пароль"})
		if err != nil {
			log.Println("Неверные данные:", err.Error())
		}
		return
	}

	token := GenerateToken(userid, "manager")
	json.NewEncoder(w).Encode(models.ActiveToken{token})
}

//AuthMerchant - company authorisation main func/ основная функция авторизации компаний
func AuthMerchant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var data models.AuthData
	json.NewDecoder(r.Body).Decode(&data)
	passed, userid := CheckMerchantData(data.Login, data.Password)
	if passed != true {
		w.WriteHeader(400)
		err := json.NewEncoder(w).Encode(models.ErrorDescr{Reason: "Неверный логин или пароль"})
		if err != nil {
			log.Println("Неверные данные:", err.Error())
		}
		return
	}

	token := GenerateToken(userid, "merchant")
	json.NewEncoder(w).Encode(models.ActiveToken{token})
}

//CheckClientData - checks client's entered data for matching DB while authorisation takes place/проверяет верность введенных клиентом данных при авторизации
func CheckClientData(login, password string) (bool, int64) {
	var clientsData []models.ClientsData
	var clientData models.ClientsData

	sqlQuery := `select login, password, id from clients_data`
	data, err := database.GetDBConn.Query(sqlQuery)
	if err != nil {
		log.Println("не удалось получить данные:", err)
	}
	for data.Next() {
		err = data.Scan(&clientData.Login, &clientData.Password, &clientData.ID)
		if err != nil {
			log.Println("Ошибка ввода данных:", err.Error())
		}
		clientsData = append(clientsData, clientData)
	}
	for _, value := range clientsData {
		if login == value.Login && clientPassCheck(password) {
			return true, value.ID
		}
	}
	return false, 0
}

//CheckManagerData - checks manager's inserted data for matching DB during authorisation/проверяет верность введенных данных при авторизации
func CheckManagerData(login, password string) (bool, int64) {
	var clientsData []models.ClientsData
	var clientData models.ClientsData

	sqlQuery := `select login, password, id from managers_data`
	data, err := database.GetDBConn.Query(sqlQuery)
	if err != nil {
		log.Println("не удалось получить данные")
	}
	for data.Next() {
		err = data.Scan(&clientData.Login, &clientData.Password, &clientData.ID)
		if err != nil {
			log.Println("Ошибка ввода данных", err.Error())
		}
		clientsData = append(clientsData, clientData)
	}
	for _, value := range clientsData {
		if login == value.Login && password == value.Password {
			return true, value.ID
		}
	}
	return false, 0
}

//CheckMerchantData - checks inserted data for matching DB during authorisation/проверяет верность введенных данных при авторизации
func CheckMerchantData(login, password string) (bool, int64) {
	var clientsData []models.MerchantsData
	var clientData models.MerchantsData

	sqlQuery := `select login, password, id from merchants_data`
	data, err := database.GetDBConn.Query(sqlQuery)
	if err != nil {
		log.Println("не удалось получить данные")
	}
	for data.Next() {
		err = data.Scan(&clientData.Login, &clientData.Password, &clientData.ID)
		if err != nil {
			log.Println("Ошибка ввода данных", err.Error())
		}
		clientsData = append(clientsData, clientData)
	}
	for _, value := range clientsData {
		if login == value.Login && merchantPassCheck(password) {
			return true, value.ID
		}
	}
	return false, 0
}

//clientPassCheck - returns list of decrypted client passwords from DB/возвращает список расшифрованных паролей клиентов из БД
func clientPassCheck (clientPass string) bool {
	var passList []string
	var pass string
	query, err := database.GetDBConn.Query(database.ClientDecryptQuery, clientPass)
	if err != nil {
		log.Println("Пароль не найден!:", err.Error())
		return false
	}
	defer query.Close()


	for query.Next() {
		err = query.Scan(&pass)
		if err != nil {
			log.Println("Ошибка расшифровки:", err.Error())
			return false
		}
		passList = append(passList, pass)
	}
	return true
}

//merchantPassCheck - returns list of decrypted company passwords from DB/возвращает список расшифрованных паролей компаний из БД
func merchantPassCheck (merchantPass string) bool {
	var passList []string
	var pass string
	query, err := database.GetDBConn.Query(database.MerchantDecryptQuery, merchantPass)
	if err != nil {
		log.Println("Ошибка расшифровки 1:", err.Error())
		return false
	}
	defer query.Close()

	for query.Next() {
		err = query.Scan(&pass)
		if err != nil {
			log.Println("Ошибка расшифровки 2:", err.Error())
			return false
		}
		passList = append(passList, pass)
	}
	return true
}

//managerPassCheck - returns list of decrypted manager passwords from DB/возвращает список расшифрованных паролей менеджеров банка из БД
func managerPassCheck (managerPass string) bool {
	var passList []string
	var pass string
	query, err := database.GetDBConn.Query(database.ManagerDecryptQuery, managerPass)
	if err != nil {
		log.Println("Ошибка расшифровки 1:", err.Error())
		return false
	}
	defer query.Close()

	for query.Next() {
		err = query.Scan(&pass)
		if err != nil {
			log.Println("Ошибка расшифровки 2:", err.Error())
			return false
		}
		passList = append(passList, pass)
	}
	return false
}

//AuthMiddleware - client authorisation middleware/моддлвар для авторизации клиентов
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Authorisation := r.Header.Get("Authorisation")
		if Authorisation == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Not Authorised"))
			return
		}
		if Authorisation[:6] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Could not Authorise"))
			return
		}
		Auth := strings.Replace(Authorisation, "Bearer ", "", 1)
		log.Println(Auth)
		userId, role, valid := ParsToken(Auth)
		if !valid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid Authorisation"))
			return
		}

		if role != "client" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid Authorisation role"))
			return
		}

		str := strconv.Itoa(userId)
		r.Header.Set("user_id", str)

		next.ServeHTTP(w, r)
	})
}

//AuthManagerMiddleware - middleware for manager authorisation/миддлавар для аворизации мнеджеров
func AuthManagerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Authorisation := r.Header.Get("Authorisation")
		if Authorisation == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Not Authorised"))
			return
		}
		if Authorisation[:6] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Could not Authorise"))
			return
		}
		Auth := strings.Replace(Authorisation, "Bearer ", "", 1)
		log.Println(Auth)
		userId, role, valid := ParsToken(Auth)
		if !valid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid Authorisation"))
			return
		}

		if role != "manager" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid Authorisation role"))
			return
		}

		str := strconv.Itoa(userId)
		r.Header.Set("user_id", str)

		next.ServeHTTP(w, r)
	})
}

//AuthMerchantMiddleware - middleware for company authorisation/миддлавар для аворизации компаний
func AuthMerchantMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Authorisation := r.Header.Get("Authorisation")
		if Authorisation == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Not Authorised"))
			return
		}
		if Authorisation[:6] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Could not Authorise"))
			return
		}
		Auth := strings.Replace(Authorisation, "Bearer ", "", 1)
		log.Println(Auth)
		userId, role, valid := ParsToken(Auth)
		if !valid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid Authorisation"))
			return
		}

		if role != "merchant" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid Authorisation role"))
			return
		}

		str := strconv.Itoa(userId)
		r.Header.Set("user_id", str)

		next.ServeHTTP(w, r)
	})
}
