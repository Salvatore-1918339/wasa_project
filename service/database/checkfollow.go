package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) Checkfollow(follower int, user_id int) (bool, error) {
	var counter int
	err := db.c.QueryRow("SELECT 1 FROM followers WHERE follower=? AND user_id=?", follower, user_id).Scan(&counter)

	// ! ERRORE NON TROVA riga CHECK
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	// true --> c'è qualcuno
	if counter != 0 {
		return true, nil
	}
	return false, nil
}
