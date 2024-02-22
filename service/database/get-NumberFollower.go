package database

func (db *appdbimpl) GetNumberFollower(user_id int) (int, error) {
	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM Followers WHERE user_id = ?", user_id).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}
