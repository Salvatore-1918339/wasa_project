package database

func (db *appdbimpl) ChangeNickname(user User, newNickname string) error {

	_, err := db.c.Exec("UPDATE Users SET nickname = ? WHERE user_id = ?", newNickname, user.User_id)
	return err
}
