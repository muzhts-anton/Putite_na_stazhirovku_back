package handlers

import (
	"cashmachine/Bank"
	"cashmachine/Utils"
	"encoding/json"
	"net/http"
)

type RespCashState struct {
	Denomination uint `json:"denomination"`
	Amount       uint `json:"amount"`
}

type RespTransactionState struct {
	Success     bool             `json:"success"`
	Transaction [6]RespCashState `json:"transaction"`
	BankStorage [6]RespCashState `json:"bankstorage"`
}

type ReqDebit struct {
	Sum uint `json:"sum"`
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var input ReqDebit
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "error bad input", http.StatusBadRequest)
		return
	}

	cashinfo, err := utils.Calculate(input.Sum)
	if err != nil {
		http.Error(w, "error bad input", http.StatusBadRequest)
		return
	}

	if bank.Bank.WithdrawMoney(cashinfo) != nil {
		http.Error(w, "insufficient funds", http.StatusBadRequest)
		return
	}

	bankstorage := bank.Bank.GetStorage()
	resp := RespTransactionState{
		true,
		[6]RespCashState{
			{utils.Banknotes[0], cashinfo[utils.Banknotes[0]]},
			{utils.Banknotes[1], cashinfo[utils.Banknotes[1]]},
			{utils.Banknotes[2], cashinfo[utils.Banknotes[2]]},
			{utils.Banknotes[3], cashinfo[utils.Banknotes[3]]},
			{utils.Banknotes[4], cashinfo[utils.Banknotes[4]]},
			{utils.Banknotes[5], cashinfo[utils.Banknotes[5]]},
		},
		[6]RespCashState{
			{utils.Banknotes[0], bankstorage[utils.Banknotes[0]]},
			{utils.Banknotes[1], bankstorage[utils.Banknotes[1]]},
			{utils.Banknotes[2], bankstorage[utils.Banknotes[2]]},
			{utils.Banknotes[3], bankstorage[utils.Banknotes[3]]},
			{utils.Banknotes[4], bankstorage[utils.Banknotes[4]]},
			{utils.Banknotes[5], bankstorage[utils.Banknotes[5]]},
		},
	}
	b, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "bruh my bad idk", http.StatusInternalServerError)
		return
	}
	w.Write(b)
}