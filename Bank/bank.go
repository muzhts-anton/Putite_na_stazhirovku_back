package bank

import (
	"cashmachine/Utils"
	"errors"
)

// TODO: add history
type MoneyVault struct {
	storage utils.Payment
}

var Bank = MoneyVault{
	utils.Payment{
		utils.Banknotes[0]: 30,
		utils.Banknotes[1]: 21,
		utils.Banknotes[2]: 14,
		utils.Banknotes[3]: 56,
		utils.Banknotes[4]: 73,
		utils.Banknotes[5]: 124,
	},
}

func (bank *MoneyVault) WithdrawMoney(request utils.Payment) error {
	if bank.storage[utils.Banknotes[0]] < request[utils.Banknotes[0]] {
		return errors.New("at 5000")
	}
	if bank.storage[utils.Banknotes[1]] < request[utils.Banknotes[1]] {
		return errors.New("at 2000")
	}
	if bank.storage[utils.Banknotes[2]] < request[utils.Banknotes[2]] {
		return errors.New("at 1000")
	}
	if bank.storage[utils.Banknotes[3]] < request[utils.Banknotes[3]] {
		return errors.New("at 500")
	}
	if bank.storage[utils.Banknotes[4]] < request[utils.Banknotes[4]] {
		return errors.New("at 200")
	}
	if bank.storage[utils.Banknotes[5]] < request[utils.Banknotes[5]] {
		return errors.New("at 100")
	}

	bank.storage[utils.Banknotes[0]] -= request[utils.Banknotes[0]]
	bank.storage[utils.Banknotes[1]] -= request[utils.Banknotes[1]]
	bank.storage[utils.Banknotes[2]] -= request[utils.Banknotes[2]]
	bank.storage[utils.Banknotes[3]] -= request[utils.Banknotes[3]]
	bank.storage[utils.Banknotes[4]] -= request[utils.Banknotes[4]]
	bank.storage[utils.Banknotes[5]] -= request[utils.Banknotes[5]]

	return nil
}

func (bank *MoneyVault) PutMoney(request utils.Payment) {
	bank.storage[utils.Banknotes[0]] += request[utils.Banknotes[0]]
	bank.storage[utils.Banknotes[1]] += request[utils.Banknotes[1]]
	bank.storage[utils.Banknotes[2]] += request[utils.Banknotes[2]]
	bank.storage[utils.Banknotes[3]] += request[utils.Banknotes[3]]
	bank.storage[utils.Banknotes[4]] += request[utils.Banknotes[4]]
	bank.storage[utils.Banknotes[5]] += request[utils.Banknotes[5]]
}

func (bank *MoneyVault) GetStorage() utils.Payment {
	return bank.storage
}
