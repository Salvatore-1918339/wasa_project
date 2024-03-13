package database

import "database/sql"

func (db *appdbimpl) CheckPhoto(user_id int, photo_id int) (bool, error) {
	var counter int
	err := db.c.QueryRow("SELECT 1 FROM Photo WHERE photo_id=? AND owner=?", photo_id, user_id).Scan(&counter)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if counter != 0 {
		return true, nil
	}
	return false, nil
}
