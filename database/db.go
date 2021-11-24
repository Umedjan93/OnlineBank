package database

import (
	"OnlineBank/settings"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx"
	"io/ioutil"
	"log"
	"os"
)

var GetDBConn *pgx.Conn

//ConnToDB - соединяемся к ДБ
func ConnToDB() {
	bytes, err := ioutil.ReadFile("Settings/settings.json")
	if err != nil {
		log.Println("Cannot read file %e", err)
		return
	}
	var Database settings.SettingsDB
	err = json.Unmarshal(bytes, &Database)
	if err != nil {
		log.Println("Переменная bytes не парсится: %e", err)
		return
	}

	dsn := pgx.ConnConfig{
		Host:                 "localhost",
		Port:                 5432,
		Database:             "BankingExamination",
		User:                 "postgres",
		Password:             "umedjana1",
		TLSConfig:            nil,
		UseFallbackTLS:       false,
		FallbackTLSConfig:    nil,
		Logger:               nil,
		LogLevel:             0,
		Dial:                 nil,
		RuntimeParams:        nil,
		OnNotice:             nil,
		CustomConnInfo:       nil,
		CustomCancel:         nil,
		PreferSimpleProtocol: false,
		TargetSessionAttrs:   "",
	}
	connection, err := pgx.Connect(dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка подключения к БД: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Вы поключены к локальной БД")
	GetDBConn = connection
}
