package database

func (db *appdbimpl) FindLikes(photo Photo_id) ([]User_id, error) {
	queryRes, err := db.c.Query("SELECT * FROM Likes WHERE photo_id=?", photo.Photo_id)
	// ! Controllo degli errori
	if err != nil {
		queryRes.Close()
		return nil, err
	}
	if queryRes.Err() != nil {
		return nil, queryRes.Err()
	}

	var users_id []User_id
	for queryRes.Next() {
		var photo Complete_Photo
		var temp_user int
		err = queryRes.Scan(&photo.Photo_id, &temp_user)
		if err != nil {
			return nil, err
		}
		users_id = append(users_id, User_id{User_id: temp_user})
	}
	return users_id, nil
}
