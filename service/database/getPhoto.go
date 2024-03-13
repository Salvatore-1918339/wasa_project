package database

func (db *appdbimpl) FindPhoto(photo_id int) (Complete_Photo, error) {
	queryRes, err := db.c.Query("SELECT Photo.photo_id, Users.user_id, Users.nickname, Photo.timestamp FROM Photo INNER JOIN Users ON Photo.owner = Users.user_id WHERE Photo.photo_id = ?", photo_id)
	var photo Complete_Photo
	// ! Controllo degli errori
	if err != nil {
		queryRes.Close()
		return photo, err
	}
	if queryRes.Err() != nil {
		return photo, queryRes.Err()
	}

	for queryRes.Next() {
		err = queryRes.Scan(&photo.Photo_id, &photo.Owner.User_id, &photo.Owner.Nickname, &photo.Timestamp) // Prendiamo i risultati ottenuti e li inseriamo nel singolo utente
		if err != nil {
			return photo, err
		}
	}

	return photo, nil
}
