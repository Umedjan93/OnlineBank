package controllers

import (
	"OnlineBank/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//ErrorMessage - структурирует ошибку, выданную через api/ sorts API errors
func ErrorMessage(w http.ResponseWriter, err error) {
	response := models.ResponseStruct{
		Message:     "Упс, появилась ошибка!",
		Error:       "1",
		Description: fmt.Sprint(err),
	}

	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		log.Println("Error Message выдала ошибку", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
