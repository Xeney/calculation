package logic

import (
	"database/sql"
	"encoding/json"
	"errors"

	// _ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
)

// Name, Description, NumberOfTickets, NetProfit, Expenses
type Event struct {
	Name        string
	Description string
	// Количество билетов
	CountTicketVip, CountTicketStandart, CountTicketEconom int
	// Цена билетов
	VipPrice, StandartPrice, EconomPrice int
	// Чистая прибыль
	NetProfit int
	// Расходы
	Expenses int
}

// Получить пользователя по login
func GetToId_sql(login string) (User, error) {
	db, err := sql.Open("sqlite", "db.sqlite")
	if err != nil {
		return User{}, errors.New("error connection")
	}
	defer db.Close()
	row := db.QueryRow("select * from users where login = $1", login)
	u := User{}
	var eventes_ []byte
	err = row.Scan(&u.Id, &u.Login, &u.Password, &eventes_)
	if err != nil {
		return User{}, errors.New("error account")
	}
	err = json.Unmarshal(eventes_, &u.Events)
	if err != nil {
		return User{}, errors.New("error unmarshal")
	}
	return u, nil
}

// Создать пользователя
func Create_sql(login, password string, eventByte map[string]Event) (sql.Result, error) {
	db, err := sql.Open("sqlite", "db.sqlite")
	if err != nil {
		return nil, errors.New("error connection")
	}
	defer db.Close()
	resEventByte, err := json.Marshal(eventByte)
	if err != nil {
		return nil, errors.New("error marshal")
	}
	result, err := db.Exec("insert into users (login, password, events) values ($1, $2, $3)", login, password, resEventByte)
	if err != nil {
		return nil, errors.New("error sql_exec")
	}
	return result, nil
}

// Добавление моих мероприятий
func UpdateFavorites_sql(countTicketVip, countTicketStandart, countTicketEconom, vipPrice, standartPrice, economPrice, NetProfit, Expenses int, Name, Description, login_user string) error {
	db, err := sql.Open("sqlite", "db.sqlite")
	if err != nil {
		return errors.New("error connection")
	}
	defer db.Close()
	user, err := GetUser(login_user)
	if err != nil {
		return err
	}
	if (user.Events[Name] != Event{}) {
		return errors.New("название занято")
	}
	user.Events[Name] = Event{Name, Description, countTicketVip, countTicketStandart, countTicketEconom, vipPrice, standartPrice, economPrice, NetProfit, Expenses}
	res_m, err := json.Marshal(user.Events)
	if err != nil {
		return errors.New("error marshal")
	}
	_, err = db.Exec("update users set events = $1 where login = $2", res_m, login_user)
	if err != nil {
		return errors.New("error exec")
	}
	return nil
}

// Удаление моих мероприятий
func DeleteFavorites_sql(Name string, login_user string) error {
	db, err := sql.Open("sqlite", "db.sqlite")
	if err != nil {
		return errors.New("error connection")
	}
	defer db.Close()
	user, err := GetUser(login_user)
	if err != nil {
		return err
	}
	if (user.Events[Name] != Event{}) {
		delete(user.Events, Name)
	} else {
		return errors.New("already in favorites")
	}
	res_m, err := json.Marshal(user.Events)
	if err != nil {
		return errors.New("error marshal")
	}
	_, err = db.Exec("update users set events = $1 where login = $2", res_m, login_user)
	if err != nil {
		return errors.New("error exec")
	}
	return nil
}
