package database

import (
	"database/sql"
)

func (db *appdbimpl) Checkfollow(follower int, user_id int) (bool, error) {
	var counter int
	err := db.c.QueryRow("SELECT 1 FROM followers WHERE follower=? AND user_id=?", follower, user_id).Scan(&counter)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	// true --> c'è qualcuno
	if counter != 0 {
		return true, nil
	}
	return false, nil
}
