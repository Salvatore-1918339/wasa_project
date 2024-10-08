package database

import "time"

func (db *appdbimpl) CreatePhoto(Owner int, timestamp time.Time) (int, error) {
	// Eseguo la query
	result, err := db.c.Exec("INSERT INTO Photo(owner,timestamp) VALUES (?,?)", Owner, timestamp)
	if err != nil {
		return -1, err
	}
	// prendo l'ID dell'ultimo valore inserito
	photoId, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(photoId), nil
}
