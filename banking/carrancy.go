package banking

import "errors"

type Currency struct {
	Symbol      string
	Description string
}

var carrencies = map[string]Currency{
	"rub": {
		Symbol:      "ru",
		Description: "Рубль",
	},
	"usd": {
		Symbol:      "us",
		Description: "Доллар",
	},
}

func GetCurrency(code string) (*Currency, error) {
	if v, ok := carrencies[code]; ok {
		return &v, nil
	} else {
		err := errors.New("Нет валюты с таким кодом! " + code)
		return &Currency{}, err
	}
}
