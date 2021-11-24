/* 	данный файл служит для инициализации таблиц в pgx и дальнейшего их создания в зависимости
от структур таблиц в файле ddl.go*/

package database

import (
	"github.com/jackc/pgx"
	"log"
)

//DbInitialisation - создаем таблицы в БД
func DbInitialisation(db *pgx.Conn) error {
	DDLs := []string{ClientsDataDDL, MerchantsDataDDL, ManagersDataDDL, ClientAccountsDDL,
		MerchantAccountsDDL, ServicesDDL, TransactionsDDL, ATMsDDL}
	for _, ddl := range DDLs {
		_, err := db.Exec(ddl)
		if err != nil {
			log.Println("cant create a table:", err)
		}
	}
	return nil
}
