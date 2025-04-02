package logic

import (
	"database/sql"
	"errors"
)

type User struct {
	Id       int
	Login    string
	Password string
	Events   map[string]Event
}

func AddUser(login, password string, eventsByte map[string]Event) (sql.Result, error) {
	result, err := Create_sql(login, password, eventsByte)
	if err != nil {
		return nil, errors.New("error sql")
	}
	return result, nil
}

func GetUser(login string) (User, error) {
	result, err := GetToId_sql(login)
	if err != nil {
		return User{}, errors.New("error sql")
	}
	return result, nil
}

// Оптимальная стоимость одного билета; Сумма с продажи всех билетов; Сумма дохода с процентов; Чистая прибыль;
func EventCalculation(vipPrice, standartPrice, economPrice, countTicketVip, countTicketStandart, countTicketEconom, plusMoney, minusMoney int, checkSave, name, description, login_user string) (int, int, int, int, error) {
	// Сумма с продажи всех билетов
	summAll := (countTicketVip * vipPrice) + (countTicketStandart * standartPrice) + (countTicketEconom * economPrice)
	// Сумма дохода с процентов
	summPrc := (summAll / 100) * plusMoney
	// Чистая прибыль
	moneyUp := (summAll + summPrc) - minusMoney
	// Оптимальная стоимость одного билета
	oneTiket := moneyUp / (countTicketEconom + countTicketStandart + countTicketVip)
	if checkSave == "save" {
		err := UpdateFavorites_sql(countTicketVip, countTicketStandart, countTicketEconom, vipPrice, standartPrice, economPrice, plusMoney, minusMoney, name, description, login_user)
		if err != nil {
			return 0, 0, 0, 0, err
		}
	}
	return oneTiket, summAll, summPrc, moneyUp, nil
}
