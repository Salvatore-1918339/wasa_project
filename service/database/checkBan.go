package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) CheckBan(banned User, user User) (bool, error) {
	var counter int
	err := db.c.QueryRow("SELECT 1 FROM Banned WHERE banner=? AND user_id=?", user.User_id, banned.User_id).Scan(&counter)

	// ! ERRORE NON TROVA riga CHECK
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return true, err
	}
	// Controllo se COUNTER è vuoto
	if counter != 0 { // se è diverso da 0 significa che abbiamo trovato una riga, quindi è bannato
		return true, nil
	}
	return false, nil
}
