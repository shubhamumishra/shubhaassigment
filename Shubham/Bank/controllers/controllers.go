package controllers

import (
	"bank/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BankAPI struct {
	Db *sql.DB
}

func (api *BankAPI) CreateBank(w http.ResponseWriter, r *http.Request) {
	var bank models.Bank
	err := json.NewDecoder(r.Body).Decode(&bank)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := api.Db.Exec("INSERT INTO banks (name, ifsc_code, branch_name) VALUES (?, ?, ?)", bank.Name, bank.IFSCCode, bank.BranchName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bank.ID = id
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bank)
}

func (api *BankAPI) ListBanks(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, name, ifsc_code, branch_name FROM banks"
	branchName := r.URL.Query().Get("branch_name")
	if branchName != "" {
		query += " WHERE branch_name = ?"
	}
	rows, err := api.Db.Query(query, branchName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	banks := make([]*models.Bank, 0)
	for rows.Next() {
		var bank models.Bank
		err := rows.Scan(&bank.ID, &bank.Name, &bank.IFSCCode, &bank.BranchName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		banks = append(banks, &bank)
	}
	json.NewEncoder(w).Encode(banks)
}

func (api *BankAPI) DeleteBank(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = api.Db.Exec("DELETE FROM banks WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (api *BankAPI) GetBank(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	row := api.Db.QueryRow("SELECT id, name, ifsc_code, branch_name FROM banks WHERE id = ?", id)
	var bank models.Bank
	err = row.Scan(&bank.ID, &bank.Name, &bank.IFSCCode, &bank.BranchName)
	if err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(w).Encode(bank)
}

func (api *BankAPI) UpdateBank(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var bank models.Bank
	err = json.NewDecoder(r.Body).Decode(&bank)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = api.Db.Exec("UPDATE banks SET name = ?, ifsc_code = ?, branch_name = ? WHERE id = ?", bank.Name, bank.IFSCCode, bank.BranchName, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bank.ID = id
	json.NewEncoder(w).Encode(bank)
}

func (api *BankAPI) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account models.Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := api.Db.Exec("INSERT INTO accounts (bank_name, branch_name, account_holder, identity_id, first_name, last_name, address, bank_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", account.BankName, account.BranchName, account.AccountHolder, account.IdentityID, account.FirstName, account.LastName, account.Address, account.BankID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	account.ID = id
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

func (api *BankAPI) GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	row := api.Db.QueryRow("SELECT id, bank_name, branch_name, account_holder, identity_id, first_name, last_name, address, bank_id FROM accounts WHERE id = ?", id)
	var account models.Account
	err = row.Scan(&account.ID, &account.BankName, &account.BranchName, &account.AccountHolder, &account.IdentityID, &account.FirstName, &account.LastName, &account.Address, &account.BankID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(w).Encode(account)
}
