package database

func (db *appdbimpl) FollowUser(follower User, user User) error {
	_, err := db.c.Exec("INSERT INTO Followers(follower, user_id) VALUES (?,?)", follower.User_id, user.User_id)
	if err != nil {
		return err
	}
	return nil
}
