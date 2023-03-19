package models

type Bank struct {
    ID         int64  `json:"id"`
    Name       string `json:"name"`
    IFSCCode   string `json:"ifsc_code"`
    BranchName string `json:"branch_name"`
}

type Account struct {
    ID              int64  `json:"id"`
    BankName        string `json:"bank_name"`
    BranchName      string `json:"branch_name"`
    AccountHolder   string `json:"account_holder"`
    IdentityID      string `json:"identity_id"`
    FirstName       string `json:"first_name"`
    LastName        string `json:"last_name"`
    Address         string `json:"address"`
    BankID          int64  `json:"bank_id"`
}
