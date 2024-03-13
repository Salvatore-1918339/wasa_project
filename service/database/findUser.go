package database

func (db *appdbimpl) FIndUser(user User) (string, error) {
	var nickname string
	err := db.c.QueryRow("SELECT nickname FROM Users WHERE user_id=?", user.User_id).Scan(&nickname)
	if err != nil {
		return "", err
	}
	return nickname, nil
}

func (db *appdbimpl) CheckUser(u User) (User, error) {
	var user User
	user.User_id = 0

	err := db.c.QueryRow("SELECT * FROM Users WHERE nickname=?", u.Nickname).Scan(&user.User_id, &user.Nickname)
	if err != nil {
		return user, err
	}
	return user, nil

}
