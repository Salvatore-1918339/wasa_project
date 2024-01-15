package database

func (db *appdbimpl) CreateComment(owner User, photo Photo_id, txt string, Timestamp string) (int, error) {
	result, err := db.c.Exec("INSERT INTO Comment (photo_id, owner, txt, timestamp) VALUES (?,?,?,?)", photo.Photo_id, owner.User_id, txt, Timestamp)
	if err != nil {
		return -1, err
	}

	// prendo l'ID dell'ultimo valore inserito
	comment_id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(comment_id), nil
}
