package database

func (db *appdbimpl) GetNickname(user_id int) (string, error) {

	var user User
	err := db.c.QueryRow("SELECT * FROM Users WHERE user_id=?", user_id).Scan(&user.User_id, &user.Nickname)
	return user.Nickname, err
}

func (db *appdbimpl) SearchNickname(nickname string) (bool, error) {

	var exist int
	err := db.c.QueryRow("SELECT EXISTS (SELECT 1 FROM Users WHERE nickname = ?)", nickname).Scan(&exist)
	if exist == 1 {
		return true, err
	}
	return false, err
}
