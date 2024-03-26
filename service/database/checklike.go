package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) Checklike(photo_id int, user_id int) (bool, error) {
	var counter int
	err := db.c.QueryRow("SELECT 1 FROM Likes WHERE photo_id=? AND user_id=?", photo_id, user_id).Scan(&counter)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false, err
	}

	if counter != 0 {
		return true, nil
	}
	return false, nil
}
