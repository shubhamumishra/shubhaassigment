package main

import (
	"bank/controllers"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutes(api *controllers.BankAPI) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/banks", api.ListBanks).Methods("GET")
	r.HandleFunc("/banks", api.CreateBank).Methods("POST")
	r.HandleFunc("/banks/{id:[0-9]+}", api.DeleteBank).Methods("DELETE")
	r.HandleFunc("/banks/{id:[0-9]+}", api.GetBank).Methods("GET")
	r.HandleFunc("/banks/{id:[0-9]+}", api.UpdateBank).Methods("PUT")
	r.HandleFunc("/accounts", api.CreateAccount).Methods("POST")
	r.HandleFunc("/accounts/{id:[0-9]+}", api.GetAccount).Methods("GET")
	return r
}

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/bank_db")
	if err != nil {
		log.Fatal(err)
	}
	api := &controllers.BankAPI{Db: db}
	router := setupRoutes(api)
	log.Fatal(http.ListenAndServe(":8000", router))
}
