package database

func (db *appdbimpl) FindLikes(photo Photo_id) ([]User, error) {
	queryRes, err := db.c.Query("SELECT Users.user_id, Users.nickname FROM Users INNER JOIN Likes ON Users.user_id = Likes.user_id WHERE Likes.photo_id =?", photo.Photo_id)
	// ! Controllo degli errori
	if err != nil {
		queryRes.Close()
		return nil, err
	}
	if queryRes.Err() != nil {
		return nil, queryRes.Err()
	}

	var users []User
	for queryRes.Next() {
		var user User
		err = queryRes.Scan(&user.User_id, &user.Nickname)
		if err != nil {
			return nil, err
		}
		users = append(users, User{User_id: user.User_id, Nickname: user.Nickname})
	}
	return users, nil
}
