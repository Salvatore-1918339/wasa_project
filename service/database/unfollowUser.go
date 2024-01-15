package database

func (db *appdbimpl) UnFollowUser(user User, follow User) error {
	_, err := db.c.Exec("DELETE FROM Followers WHERE follower=? AND user_id=?", user.User_id, follow.User_id)
	if err != nil {
		return err
	}
	return nil
}
