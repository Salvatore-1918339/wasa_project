package database

import "fmt"

// GetName is an example that shows you how to query data
func (db *appdbimpl) GetNickname(user_id int) (User, error) {

	var user User
	err := db.c.QueryRow("SELECT * FROM Users WHERE user_id=?", user_id).Scan(&user.User_id, &user.Nickname)
	fmt.Print("Error: ", err)
	return user, nil
}
