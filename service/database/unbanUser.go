package database

func (db *appdbimpl) UnBanUser(banner int, user_id int) error {

	_, err := db.c.Exec("DELETE FROM Banned WHERE banner=? AND user_id=?", banner, user_id)
	if err != nil {
		return err
	}
	return nil
}
