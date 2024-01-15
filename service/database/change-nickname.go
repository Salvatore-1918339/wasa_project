package database

import "fmt"

func (db *appdbimpl) ChangeNickname(user User, newNickname string) error {

	_, err := db.c.Exec("UPDATE Users SET nickname = ? WHERE user_id = ?", newNickname, user.User_id)
	if err != nil {
		fmt.Print("#!# Errore in ChangeNickname API: ", err)
	}
	fmt.Print(newNickname, " : Nickname Cambiato Correttamente.\n")
	return nil
}
