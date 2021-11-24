package repositories

import (
	"OnlineBank/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//StartRoutes - запускает роут/starts rout
func StartRoutes() {

	r := mux.NewRouter()
	log.Println("Слушаем порт 8888...")

	client := r.PathPrefix("/").Subrouter()
	client.Use(controllers.AuthMiddleware)
	manager := r.PathPrefix("/").Subrouter()
	manager.Use(controllers.AuthManagerMiddleware)
	merchant := r.PathPrefix("/").Subrouter()
	merchant.Use(controllers.AuthMerchantMiddleware)

	r.HandleFunc("/client/auth", controllers.AuthClient)
	r.HandleFunc("/merchant/auth", controllers.AuthMerchant)
	r.HandleFunc("/manager/auth", controllers.AuthManager)

	r.HandleFunc("/main/registration", controllers.RegistrationHandler).Methods("POST")

	client.HandleFunc("/client/services", controllers.GetServicesHandler).Methods("GET")
	client.HandleFunc("/client/ATMs", controllers.GetAtmsHandler).Methods("GET")
	client.HandleFunc("/client/transferBtwClients", controllers.TxAccHandler).Methods("POST")
	client.HandleFunc("/client/transferByPhone", controllers.TxPhoneHandler).Methods("POST")
	client.HandleFunc("/client/serviceTx", controllers.ServiceTxHandler).Methods("POST")

	merchant.HandleFunc("/merchant/newService", controllers.AddServiceHandler).Methods("POST")
	merchant.HandleFunc("/merchant/updateService", controllers.UpdateServiceHandler).Methods("PUT")
	merchant.HandleFunc("/merchant/deleteService", controllers.DeleteServiceHandler).Methods("POST")

	manager.HandleFunc("/manager/add-data/addClient", controllers.AddClientHandler).Methods("POST")
	manager.HandleFunc("/manager/add-data/addMerchant", controllers.AddMerchantHandler).Methods("POST")
	manager.HandleFunc("/manager/add-data/addMerchantAcc", controllers.AddMerchantAccHandler).Methods("POST")
	manager.HandleFunc("/manager/add-data/addClientAcc", controllers.AddClientAccHandler).Methods("POST")
	manager.HandleFunc("/manager/add-data/addService", controllers.AddServiceHandler).Methods("POST")
	manager.HandleFunc("/manager/add-data/addAtm", controllers.AddAtmHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8888", r))
}
