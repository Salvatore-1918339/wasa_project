package database

func (db *appdbimpl) GetNumberFollowed(user_id int) (int, error) {
	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM Followers WHERE follower = ?", user_id).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}
