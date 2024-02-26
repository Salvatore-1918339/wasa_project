package database

func (db *appdbimpl) FindPhotos(user User) ([]Complete_Photo, error) {
	queryRes, err := db.c.Query("SELECT photo_id, timestamp FROM Photo WHERE owner=?", user.User_id)
	if err != nil {
		return nil, err
	}

	var photos []Complete_Photo

	for queryRes.Next() {

		// ! Prendiamo le info dal DB delle Photo
		var photo Complete_Photo
		photo.Owner = user.User_id
		err = queryRes.Scan(&photo.Photo_id, &photo.Timestamp) // Prendiamo i risultati ottenuti e li inseriamo nel singolo utente
		if err != nil {
			return nil, err
		}

		// ! Prendiamo i commenti della photo
		var comments []Comment
		comments, err := db.FindComment(Photo_id{Photo_id: photo.Photo_id})
		if err != nil {
			return nil, err
		}
		photo.Comments = append(photo.Comments, comments...)

		// ! Prendiamo i Like
		var likes []User_id
		likes, err = db.FindLikes(Photo_id{Photo_id: photo.Photo_id})
		if err != nil {
			return nil, err
		}
		photo.Likes = append(photo.Likes, likes...)
		photos = append(photos, photo)
	}
	return photos, nil
}
