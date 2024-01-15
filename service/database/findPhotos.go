package database

func (db *appdbimpl) FindPhotos(user User) ([]Complete_Photo, error) {
	queryRes, err := db.c.Query("SELECT photo_id, timestamp FROM Photo WHERE owner=?", user.User_id)
	if err != nil {
		return nil, err
	}

	var photos []Complete_Photo // Creo un'array di Utenti
	for queryRes.Next() {
		var photo Complete_Photo
		photo.Owner = user.User_id
		err = queryRes.Scan(&photo.Photo_id, &photo.Timestamp) // Prendiamo i risultati ottenuti e li inseriamo nel singolo utente
		if err != nil {
			return nil, err
		}
		photos = append(photos, photo) // Appendo gli User
	}
	return photos, nil
}
