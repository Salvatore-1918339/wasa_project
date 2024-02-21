package database

import (
	"fmt"
	"strconv"
)

//Metodo per aggiungere un utente nel DB
func (db *appdbimpl) CreateUser(u User) error {

	// Eseguiamo una INSERT nel DB
	result, err := db.c.Exec("INSERT INTO Users(nickname) VALUES (?)", u.Nickname)

	// Gestione degli errori
	if err != nil {
		return err
	}

	UserId, err := result.LastInsertId()
	fmt.Print("\n Id Assegnato nel DB : ", UserId)
	return nil

}

func (db *appdbimpl) FindUserId(u User) (int, error) {
	var id string
	err := db.c.QueryRow("SELECT user_id FROM Users WHERE nickname=?", u.Nickname).Scan(&id)
	if err != nil {
		return -1, err
	}
	int_id, _ := strconv.Atoi(id)
	return int_id, nil
}
