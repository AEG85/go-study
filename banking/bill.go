package banking

import "errors"

type Bill struct {
	Carrancy Currency
	Amount   int
}

func NewBill(carrency *Currency, amount int) (*Bill, error) {
	if carrency == nil {
		err := errors.New("Указан не верный формат валюты!")
		return &Bill{}, err
	}
	if amount < 0 {
		err := errors.New("Нельзя указать отрицательный баланс!")
		return &Bill{}, err
	}
	return &Bill{Carrancy: *carrency, Amount: amount}, nil
}

func (b *Bill) TakeCach(amount int) (int, error) {
	if b.Amount < amount {
		err := errors.New("Недостаточно средств на счете!")
		return b.Amount, err
	}
	b.Amount -= amount
	return b.Amount, nil
}

func (b *Bill) Pay(coast int) (int, error) {
	if b.Amount < coast {
		err := errors.New("Недостаточно средств на счете!")
		return b.Amount, err
	}
	b.Amount -= coast
	return b.Amount, nil
}

func (b *Bill) GetBalance() int {
	return b.Amount
}
